package initialize

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-boot-starter/config"
	"go-boot-starter/logger"
	"go-boot-starter/middleware"
	"go-boot-starter/router"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func StartHttpServer(cfg *config.AppConfig) error {
	server := &http.Server{
		Addr:    ":" + cfg.HttpPort,
		Handler: initEngine(cfg),
	}
	ctx, cancel := context.WithCancel(context.Background())

	// 优雅的关闭线程
	go listenToSystemSignals(cancel)
	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			logger.Error(fmt.Sprintf("Failed to shutdown server: %s", err))
		}
	}()

	addrList, err := net.InterfaceAddrs()
	if err != nil {
		logger.Debug("获取本地ip异常: " + err.Error())
		os.Exit(1)
	}

	for _, address := range addrList {
		if aspnet, ok := address.(*net.IPNet); ok && !aspnet.IP.IsLoopback() {
			if aspnet.IP.To4() != nil {
				logger.Debug(fmt.Sprintf("Server started success: http://%s:%s", aspnet.IP.String(), cfg.HttpPort))
			}
		}
	}

	if err = server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Debug("Server was shutdown gracefully")
			return nil
		}
		return err
	}
	return nil
}

func initEngine(cfg *config.AppConfig) *gin.Engine {
	gin.SetMode(func() string {
		if cfg.IsDevEnv() {
			return gin.DebugMode
		}
		return gin.ReleaseMode
	}())
	engine := gin.Default()
	engine.Use(middleware.Cors(), gin.Recovery())

	//engine.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
	//	c.AbortWithStatusJSON(http.StatusOK, gin.H{
	//		"code": 500,
	//		"msg":  "服务器内部错误，请稍后再试！",
	//	})
	//}))
	router.RegisterRoutes(engine)
	return engine
}

func listenToSystemSignals(cancel context.CancelFunc) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case <-signalChan:
			cancel()
			return
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"go-boot-starter/config"
	"go-boot-starter/initialize"
	"go-boot-starter/logger"
	"go-boot-starter/models"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "conf", "app.ini", "select active profile")
	flag.Parse()

	// 加载配置
	cfg, err := config.Load(fmt.Sprintf("conf/%s", configFile))
	if err != nil {
		panic(fmt.Sprintf("load config failed, file: %s, error: %s", configFile, err))
	}

	// 初始化日志
	logger.Init(cfg)
	logger.Info(fmt.Sprintf("\n %+v", cfg))

	if err := models.InitDB(cfg); err != nil {
		panic(fmt.Sprintf("db start error: %s", err))
	}

	//启动http服务器
	if err := initialize.StartHttpServer(cfg); err != nil {
		panic(fmt.Sprintf("server start error: %s", err))
	}

}

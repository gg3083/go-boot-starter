package router

import (
	"github.com/gin-gonic/gin"
	"go-boot-starter/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", controllers.BaseController.Index)
	server.GET("/", controllers.NovelControllers.Index)
	//r := server.Group("/api")

}

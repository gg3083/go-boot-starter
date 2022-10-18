package router

import (
	"github.com/gin-gonic/gin"
	"go-boot-starter/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", controllers.BaseController.Index)
	r := server.Group("/api/v1")
	novel := r.Group("/novel")
	{
		novel.GET("/", controllers.NovelControllers.Index)
		novel.GET("/index", controllers.NovelControllers.NovelIndexList)
		novel.GET("/details", controllers.NovelControllers.NovelIndexDetail)
		novel.GET("/content", controllers.NovelControllers.NovelContentList)
		novel.GET("/content/details", controllers.NovelControllers.NovelIndexDetail)
	}

}

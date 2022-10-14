package controllers

import (
	"github.com/gin-gonic/gin"
	"go-boot-starter/pkg/app"
)

type NovelController struct{}

var NovelControllers = &NovelController{}

func (*NovelController) Index(ctx *gin.Context) {
	app.SuccessResp(ctx)
}

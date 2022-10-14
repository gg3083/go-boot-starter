package controllers

import (
	"github.com/gin-gonic/gin"
	"go-boot-starter/pkg/app"
)

type Controller struct{}

var BaseController = &Controller{}

func (*Controller) Index(ctx *gin.Context) {
	app.SuccessResp(ctx)
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"go-boot-starter/pkg/app"
	"go-boot-starter/service"
)

type NovelController struct {
}

var NovelControllers = &NovelController{}

func (*NovelController) Index(ctx *gin.Context) {
	app.SuccessResp(ctx)
}

type NovelIndexForm struct {
	PageNo    int    `json:"pageNo" form:"pageNo"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
	SearchKey string `json:"searchKey" form:"searchKey"`
}

func (*NovelController) NovelIndexList(ctx *gin.Context) {
	var form NovelIndexForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		app.ParamErrorResp(ctx, "参数不正确")
		return
	}
	novel_service := &service.NovelService{
		PageNo:   form.PageNo,
		PageSize: form.PageSize,
	}
	total, indices, err := novel_service.ListNovelIndex()
	if err != nil {
		app.ErrorResp(ctx, err)
		return
	}
	app.SuccessRespData(ctx, &struct {
		Total int         `json:"total"`
		List  interface{} `json:"list"`
	}{
		Total: total,
		List:  indices,
	})
}

type NovelIndexDetailForm struct {
	Id int `json:"id" form:"id"`
}

func (*NovelController) NovelIndexDetail(ctx *gin.Context) {
	var form NovelIndexDetailForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		app.ParamErrorResp(ctx, "参数不正确")
		return
	}
	novel_service := &service.NovelService{
		Id: form.Id,
	}
	detail, err := novel_service.GetNovelIndex()
	if err != nil {
		app.ErrorResp(ctx, err)
		return
	}
	app.SuccessRespData(ctx, &struct {
		Data interface{} `json:"data"`
	}{
		Data: detail,
	})
}

type NovelDetailForm struct {
	PageNo    int    `json:"pageNo" form:"pageNo"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
	SearchKey string `json:"searchKey" form:"searchKey"`
}

func (*NovelController) NovelContentList(ctx *gin.Context) {
	var form NovelDetailForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		app.ParamErrorResp(ctx, "参数不正确")
		return
	}
	novel_service := &service.NovelService{
		//Id: form.Id,
	}
	total, detail, err := novel_service.ListNovelContent()
	if err != nil {
		app.ErrorResp(ctx, err)
		return
	}
	app.SuccessRespData(ctx, &struct {
		Data  interface{} `json:"data"`
		Total int         `json:"total"`
	}{
		Data:  detail,
		Total: total,
	})
}

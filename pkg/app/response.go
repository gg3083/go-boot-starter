package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	Success      = 0   //正常
	Failed       = 1   //失败
	ParamError   = 400 //参数错误
	Unauthorized = 401 //未认证
	NotFound     = 404 //不存在
	UnAuthorized = 401 //未授权
	NotLogin     = 405 //未登录
)

// ErrorRespByCode ErrorResp 错误返回值
func ErrorRespByCode(c *gin.Context, code int, msg string) {
	respOk(c, code, msg, nil)
}

// ErrorResp 错误返回值
func ErrorResp(c *gin.Context, err error) {
	respOk(c, Failed, err.Error(), nil)
}

// ParamErrorResp ErrorResp 错误返回值
func ParamErrorResp(c *gin.Context, errMsg string) {
	respOk(c, ParamError, errMsg, nil)
}

// UnauthorizedResp ErrorResp 错误返回值
func UnauthorizedResp(c *gin.Context, errMsg string) {
	respOk(c, Unauthorized, errMsg, nil)
}

// SuccessRespByCode SuccessResp 正确返回值
func SuccessRespByCode(c *gin.Context, code int, data interface{}) {
	respOk(c, code, "", data)
}

// SuccessRespData 正确返回值
func SuccessRespData(c *gin.Context, data interface{}) {
	respOk(c, Success, "", data)
}

// SuccessResp 正确返回值
func SuccessResp(c *gin.Context) {
	respOk(c, Success, "", nil)
}

// resp 返回
func respOk(c *gin.Context, code int, msg string, data interface{}) {
	respByHttpCode(c, http.StatusOK, code, msg, data)
}

// resp 返回
func respByHttpCode(c *gin.Context, httpCode, code int, msg string, data interface{}) {
	c.JSON(httpCode, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

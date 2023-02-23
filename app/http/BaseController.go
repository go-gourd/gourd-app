package http

import (
	"github.com/gin-gonic/gin"
)

// BaseController 基础控制器
type BaseController struct {
}

// Success 成功时响应
func (*BaseController) Success(c *gin.Context, message string, data any) {
	if message == "" {
		message = "success"
	}
	res := gin.H{
		"code":    0,
		"data":    data,
		"message": message,
	}
	c.AsciiJSON(200, &res)
}

// Fail 失败响应
func (*BaseController) Fail(c *gin.Context, code int, message string, data any) {
	if message == "" {
		message = "fail"
	}
	res := gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	}
	c.AsciiJSON(200, &res)
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/log"
)

func RegisterMiddleware(app *gin.Engine) {

	log.Info("注册中间件")

	//示例中间件
	app.Use(ExampleMiddleware())

}

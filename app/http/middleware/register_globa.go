package middleware

import (
	"github.com/gin-gonic/gin"
)

// RegisterGlobalMiddleware 注册全局中间件
func RegisterGlobalMiddleware(app *gin.Engine) {

	//示例中间件
	//app.Use(ExampleMiddleware())

	app.Use(SessionMiddle())

}

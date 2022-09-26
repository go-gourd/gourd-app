package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("ExampleMiddleware")

		//日志输出请求Header的Accept
		//log.Info(c.GetHeader("Accept"))

	}
}

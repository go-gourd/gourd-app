package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// ExampleMiddleware 示例中间件
func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Print("User-Agent: ")
		fmt.Println(c.GetHeader("User-Agent"))

	}
}

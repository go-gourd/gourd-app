package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func ExampleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Print(c.GetHeader("Accept"))

	}
}

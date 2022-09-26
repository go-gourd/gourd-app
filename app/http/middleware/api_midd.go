package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ApiBaseMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Print("ApiPath: ")
		fmt.Println(c.FullPath())

	}
}

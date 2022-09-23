package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRouterGroup(router *gin.RouterGroup) {

	test := router.Group("/test")
	{
		test.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello test")
		})
	}
}

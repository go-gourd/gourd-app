package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterIndexRoute(router *gin.Engine) {

	v1 := router.Group("/v1")
	{

		qq := v1.Group("/qq")
		{
			qq.GET("/login", func(c *gin.Context) {
				c.String(http.StatusOK, "Hello login")
			})
		}
	}
}

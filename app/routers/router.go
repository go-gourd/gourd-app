package routers

import (
	apiR "app/routers/api"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	api := r.Group("/api")
	{
		apiR.AddUserRouter(api)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

package index

import (
	"github.com/gin-gonic/gin"
	"gourd/app/http/controllers/index"
)

func RegisterIndexRoute(router *gin.Engine) {

	router.GET("/", index.Controller{}.Index)
	router.GET("/session", index.Controller{}.Session)

}

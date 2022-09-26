package api

import (
	"github.com/gin-gonic/gin"
	"gourd/app/http/controllers/api/test"
)

func RegisterRouterGroup(router *gin.RouterGroup) {

	router.GET("/test/getDate", test.Date{}.Get)

}

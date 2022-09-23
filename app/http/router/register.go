package router

import (
	"ginapp/app/http/router/api"
	"ginapp/app/http/router/index"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {

	//注册Index路由组
	index.RegisterIndexRoute(router)

	//注册Api路由组
	apiRouter := router.Group("/api")
	{
		api.RegisterRouterGroup(apiRouter)
	}
}

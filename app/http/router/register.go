package router

import (
	"github.com/gin-gonic/gin"
	"gourd/app/http/router/api"
	"gourd/app/http/router/index"
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

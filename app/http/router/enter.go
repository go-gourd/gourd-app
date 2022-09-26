package router

import (
	"github.com/gin-gonic/gin"
	"gourd/app/http/middleware"
	"gourd/app/http/router/api"
	"gourd/app/http/router/index"
)

// RegisterRoute 路由注册入口
func RegisterRoute(router *gin.Engine) {

	//注册Index路由
	index.RegisterIndexRoute(router)

	//注册Api路由组
	apiRouter := router.Group("/api")
	{
		//添加Api中间件
		apiRouter.Use(middleware.ApiBaseMiddle())

		api.RegisterRouterGroup(apiRouter)
	}
}

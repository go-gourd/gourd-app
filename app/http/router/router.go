package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/ghttp"
	apiRoute "gourd/app/http/api/route"
)

// RegisterRouter 初始化路由
func RegisterRouter() {

	r := ghttp.GetEngine()

	//主页
	r.Any("/", func(c *gin.Context) {
		ghttp.Write(c, "hello gourd!")
	})

	//404
	r.NoRoute(func(c *gin.Context) {
		ghttp.Write(c, "404 not found")
	})

	//注册api相关路由
	apiGroup := r.Group("/api")
	{
		apiRoute.RegisterRoute(apiGroup)
	}

}

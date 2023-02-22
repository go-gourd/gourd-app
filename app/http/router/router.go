package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/ghttp"
	"gourd/app/http/router/api"
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
	apiG := r.Group("/api")
	{
		api.RegisterRoute(apiG)
	}

}

package main

import (
	"app/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	//初始化路由
	routers.InitRouter(r)

	// listen and serve
	err := r.Run()
	if err != nil {
		panic(err)
	}
}

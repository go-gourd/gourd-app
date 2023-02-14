package main

import (
	"app/routers"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/config"
	"strconv"
)

func main() {

	fmt.Println("starting...")

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	//初始化路由
	routers.InitRouter(r)

	appConfig := config.GetAppConfig()

	listen := appConfig.Host + ":" + strconv.Itoa(int(appConfig.Port))

	// listen and serve
	err := r.Run(listen)
	if err != nil {
		panic(err)
	}
}

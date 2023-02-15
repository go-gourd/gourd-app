package main

import (
	"app/http/router"
	"github.com/go-gourd/gourd"
)

func main() {

	app := gourd.App{}

	r := app.Create()

	//初始化路由
	router.InitRouter(r)

	app.Run()

}

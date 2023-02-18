package main

import (
	"github.com/go-gourd/gourd"
	"gourd/app/event"
	"gourd/app/http/router"
)

func main() {

	app := gourd.App{}

	//注册事件
	event.RegisterEvent()

	app.Init()

	r := app.Create()

	//初始化路由
	router.InitRouter(r)

	app.Run()

}

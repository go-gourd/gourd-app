package main

import (
	"github.com/go-gourd/gourd"
	"gourd/app/event"
)

func main() {

	// 实例化app
	app := gourd.App{}

	// 注册事件
	event.RegisterEvent()

	// 执行初始化
	app.Init()

	// 启动应用
	app.Run()

}

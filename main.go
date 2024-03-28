package main

import (
	"github.com/go-gourd/gourd"
	"gourd/internal/event"
)

func main() {

	// 实例化app
	app := gourd.App{}

	// 注册事件
	event.Register()

	// 执行初始化
	app.Init()

	// 启动应用
	app.Run()

}

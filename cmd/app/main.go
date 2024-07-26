package main

import (
	"github.com/go-gourd/gourd"
	"gourd/internal/event"
)

func main() {

	// 创建一个应用实例
	app := gourd.App{}

	// 注册事件
	event.Register(&app)

	// 执行初始化
	app.Init()

	// 启动应用
	app.Run()

}

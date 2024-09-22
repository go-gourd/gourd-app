package main

import (
	"context"
	"github.com/go-gourd/gourd"
	"gourd/internal/event"
)

func main() {

	ctx := context.WithValue(context.Background(), "cmd", "app")

	// 创建一个应用实例
	app := gourd.App{
		// 应用事件初始化入口
		EventHandler: event.Register,
		Context:      ctx,
	}

	// 执行初始化
	app.Init()

	// 启动应用
	app.Run()

}

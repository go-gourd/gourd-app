package main

import (
	"app/internal/event"
	"context"

	"github.com/go-gourd/gourd"
)

func main() {

	ctx := context.WithValue(context.Background(), "app_name", "app")

	// 创建一个应用实例
	app := gourd.App{
		// 应用事件初始化入口
		EventHandler: event.AppEvent,
		Context:      ctx,
	}

	// 执行初始化
	app.Init()

	// 启动应用
	app.Run()

}

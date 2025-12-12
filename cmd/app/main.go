package main

import (
	"app/internal/event"
	"context"

	"github.com/go-gourd/gourd"
)

func main() {

	ctx := context.Background()

	// 创建一个应用实例
	app := gourd.App{
		EventHandler: event.AppEvent, // 应用事件初始化
		Context:      ctx,
	}

	// 执行初始化
	app.Init()

	// 启动应用
	app.Run()

}

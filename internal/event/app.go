package event

import (
	"context"
	"github.com/go-gourd/gourd/event"
	"gourd/internal/initialize"
	"log/slog"
)

// Register 事件注册
func Register(_ context.Context) {

	// Boot事件(应用) -初始化应用时执行
	event.Listen("app.boot", func(ctx context.Context) {
		slog.Debug("boot event.")

		err := initialize.InitLog()
		if err != nil {
			panic(err)
		}

		err = initialize.InitDatabase()
		if err != nil {
			panic(err)
		}

		// 初始化并执行命令行
		initialize.InitCmd()
	})

	// Init事件(应用) -初始化完成执行
	event.Listen("app.init", func(context.Context) {
		slog.Debug("init event.")

		// 初始化定时任务
		initialize.InitCron()

		// 初始化Http服务
		initialize.InitHttpServer()
	})

	// Start事件(应用) -启动后执行
	event.Listen("app.start", func(context.Context) {
		slog.Debug("start event.")

	})

	// Stop事件(应用) -终止时执行
	event.Listen("app.stop", func(context.Context) {
		slog.Debug("stop event.")
	})

	// 注册更多自定义事件监听

}

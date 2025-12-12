package event

import (
	"context"
	"log/slog"

	"github.com/go-gourd/gourd/event"

	"app/internal/initialize"
)

// AppEvent 事件注册
func AppEvent(_ context.Context) {

	// Boot事件(应用) -初始化应用时执行
	event.Listen("app.boot", func(ctx context.Context) {
		slog.Debug("boot event.")

		err := initialize.InitLog()
		if err != nil {
			panic(err)
		}

		//err = initialize.InitDatabase()
		//if err != nil {
		//    panic(err)
		//}
	})

	// Init事件(应用) -初始化完成执行
	event.Listen("app.init", func(context.Context) {
		slog.Debug("init event.")

		// 初始化命令行并解析参数
		initialize.InitCmd()
	})

	// Start事件(应用) -启动后执行
	event.Listen("app.start", func(context.Context) {
		slog.Debug("start event.")

		// 初始化定时任务
		initialize.InitCron()

		// 初始化Http服务
		initialize.InitHttpServer()

	})

	// Stop事件(应用) -停止时执行
	event.Listen("app.stop", func(context.Context) {
		slog.Debug("stop event.")
	})

}

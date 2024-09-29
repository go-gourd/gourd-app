package event

import (
	"context"
	"github.com/go-gourd/gourd/event"
	"gourd/internal/cmd"
	"gourd/internal/cron"
	"gourd/internal/http/router"
	"gourd/internal/util"
	"log/slog"
)

// RegisterAppEvent 事件注册
func RegisterAppEvent(_ context.Context) {

	// Boot事件(应用) -初始化应用时执行
	event.Listen("app.boot", func(context.Context) {
		slog.Debug("boot event.")

		err := util.InitLog()
		if err != nil {
			panic(err)
		}

		err = util.InitDatabase()
		if err != nil {
			panic(err)
		}

		// 注册命令行
		cmd.Register()
	})

	// Init事件(应用) -初始化完成执行
	event.Listen("app.init", func(context.Context) {
		slog.Debug("init event.")

		// 注册定时任务
		cron.Register()

		// 注册路由
		router.Register()
	})

	// Start事件(应用) -启动后执行
	event.Listen("app.start", func(context.Context) {
		slog.Debug("start event.")

		// 启动Http服务
		router.StartServer()
	})

	// Stop事件(应用) -终止时执行
	event.Listen("app.stop", func(context.Context) {
		slog.Debug("stop event.")
	})

	// 注册更多自定义事件监听

}

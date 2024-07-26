package event

import (
	"github.com/go-gourd/gourd"
	"github.com/go-gourd/gourd/event"
	"gourd/internal/cmd"
	"gourd/internal/cron"
	"gourd/internal/http/router"
	"log/slog"
)

// Register 事件注册
func Register(app *gourd.App) {

	// Boot事件(应用) -初始化应用时执行
	event.Listen("app.boot", func(params any) {
		slog.Debug("boot event.")

		// TODO: 连接数据库
		//err := tools.InitDatabase()
		//if err != nil {
		//	panic(err.Error())
		//}

		// 注册命令行
		cmd.Register()
	})

	// Init事件(应用) -初始化完成执行
	event.Listen("app.init", func(params any) {
		slog.Debug("init event.")

		// 注册定时任务
		cron.Register()

		// 注册路由
		router.Register()
	})

	// Start事件(应用) -启动后执行
	event.Listen("app.start", func(params any) {
		slog.Debug("start event.")

		// 启动Http服务
		router.StartServer()
	})

	// Stop事件(应用) -终止时执行
	event.Listen("app.stop", func(ctx any) {
		slog.Debug("stop event.")
	})

	// 注册更多自定义事件监听

}

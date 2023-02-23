package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/gdb"
	"github.com/go-gourd/gourd/logger"
	"gourd/app/cmd"
	"gourd/app/dal/query"
	"gourd/app/http/router"
)

// RegisterEvent 事件注册
func RegisterEvent() {

	// Boot事件(框架) -启动时执行
	event.AddEvent("_boot", func(params any) {
		logger.Debug("boot event.")

		// 注册命令行
		cmd.RegisterCmd()
	})

	// Init事件(框架) -初始化完成执行
	event.AddEvent("_init", func(params any) {
		logger.Debug("init event.")

		// 注册路由
		router.RegisterRouter()

		// 连接数据库
		query.SetDefault(gdb.GetMysqlDb())

	})

	// Start事件(框架) -启动后执行
	event.AddEvent("_start", func(params any) {
		logger.Debug("start event.")
	})

	// Stop事件(框架) -终止时执行
	event.AddEvent("_stop", func(params any) {
		logger.Debug("stop event.")
	})

}

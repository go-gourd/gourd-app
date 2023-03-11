package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/logger"
	"github.com/go-gourd/mysql"
	"gourd/app/cmd"
	"gourd/app/cron"
	"gourd/app/dal/query"
	"gourd/app/http/router"
)

// RegisterEvent 事件注册
func RegisterEvent() {

	// Boot事件(系统) -启动时执行
	event.AddEvent("_boot", func(params any) {
		logger.Debug("boot event.")

		// 注册命令行
		cmd.RegisterCmd()

		// 注册定时任务
		cron.RegisterCron()
	})

	// Init事件(系统) -初始化完成执行
	event.AddEvent("_init", func(params any) {
		logger.Debug("init event.")

		// 注册路由
		router.RegisterRouter()

		// 连接数据库
		dbMysql, err := mysql.GetDb("mysql")
		if err != nil {
			panic(err.Error())
		}
		query.SetDefault(dbMysql)

	})

	// Start事件(系统) -启动后执行
	event.AddEvent("_start", func(params any) {
		logger.Debug("start event.")
	})

	// Stop事件(系统) -终止时执行
	event.AddEvent("_stop", func(params any) {
		logger.Debug("stop event.")
	})

}

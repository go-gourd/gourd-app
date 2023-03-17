package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/logger"
	"github.com/go-gourd/mysql"
	"gourd/app/cmd"
	"gourd/app/cron"
	"gourd/app/http/router"
	"gourd/app/orm/query"
)

// Register 事件注册
func Register() {

	// Boot事件(系统) -启动时执行
	event.Listen("_boot", func(params any) {
		logger.Debug("boot event.")

		// 注册命令行
		cmd.RegisterCmd()

		// 注册定时任务
		cron.RegisterCron()
	})

	// Init事件(系统) -初始化完成执行
	event.Listen("_init", func(params any) {
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
	event.Listen("_start", func(params any) {
		logger.Debug("start event.")
	})

	// Stop事件(系统) -终止时执行
	event.Listen("_stop", func(params any) {
		logger.Debug("stop event.")
	})

}

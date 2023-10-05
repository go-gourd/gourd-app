package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/log"
	"gourd/app/cmd"
	"gourd/app/cron"
	"gourd/app/http"
	"gourd/app/utils"
)

// Register 事件注册
func Register() {

	// Boot事件(应用) -初始化应用时执行
	event.Listen("app.boot", func(params any) {
		log.Debug("boot event.")

		// 连接数据库
		err := utils.InitDatabase()
		if err != nil {
			panic(err.Error())
		}

		// 注册命令行
		cmd.RegisterCmd()

	})

	// Init事件(应用) -初始化完成执行
	event.Listen("app.init", func(params any) {
		log.Debug("init event.")

		// 注册定时任务
		cron.Register()

		// 注册路由
		http.RegisterRouter()

	})

	// Start事件(应用) -启动后执行
	event.Listen("app.start", func(params any) {
		log.Debug("start event.")

		// 启动Http服务
		http.Start()
	})

	// Stop事件(应用) -终止时执行
	event.Listen("app.stop", func(params any) {
		log.Debug("stop event.")
	})

	// 测试事件
	event.Listen("test", func(params any) {
		log.Debug("test.")
	})

}

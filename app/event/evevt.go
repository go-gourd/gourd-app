package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/gdb"
	"github.com/go-gourd/gourd/log"
	"gourd/app/dal/query"
	"gourd/app/http/router"
)

// RegisterEvent 事件注册
func RegisterEvent() {

	// Boot事件(框架) -启动时执行
	event.AddEvent("_boot", func(params any) {
		log.Debug("boot event.")
	})

	// Init事件(框架) -初始化完成执行
	event.AddEvent("_init", func(params any) {
		log.Debug("init event.")

		//连接数据库
		query.SetDefault(gdb.GetMysqlDb())

		//注册路由
		router.InitRouter()

	})

	// Start事件(框架) -启动后执行
	event.AddEvent("_start", func(params any) {
		log.Debug("start event.")
	})

	// Stop事件(框架) -终止时执行
	event.AddEvent("_stop", func(params any) {
		log.Debug("stop event.")
	})

}

package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/log"
)

// RegisterEvent 事件注册
func RegisterEvent() {

	// Boot事件(框架) -启动时执行
	event.AddEvent("_boot", func(params any) {
		log.Debug("boot event.")
	})

	// Init事件(框架) -初始化完成执行
	event.AddEvent("_init", func(params any) {
		log.Debug("boot init.")

		//连接数据库
		//query.SetDefault(gdb.GetMysqlDb())

	})

	// Start事件(框架) -启动后执行
	event.AddEvent("_start", func(params any) {
		log.Debug("boot start.")
	})

	// Stop事件(框架) -终止时执行
	event.AddEvent("_stop", func(params any) {
		log.Debug("boot stop.")
	})

}

package event

import "github.com/go-gourd/gourd"

// RegisterEvent 注册事件
func RegisterEvent() {

	//系统全局事件
	gourd.RegisterEvent("boot", BootEvent)
	gourd.RegisterEvent("init", InitEvent)
	gourd.RegisterEvent("init", StartEvent)

}

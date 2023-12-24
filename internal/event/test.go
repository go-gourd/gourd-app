package event

import (
	"github.com/go-gourd/gourd/event"
	"github.com/go-gourd/gourd/log"
)

// 测试类事件注册
func testRegister() {

	// 测试事件
	event.Listen("test", func(params any) {
		log.Debug("Test event.")
	})

}

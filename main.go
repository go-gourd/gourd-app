package main

import (
	"ginapp/app/event"
	"github.com/go-gourd/gourd"
)

func main() {

	// 注册系统事件
	event.RegisterEvent()

	// 启动服务
	err := gourd.StartServer()
	if err != nil {
		return
	}

}

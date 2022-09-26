package main

import (
	"github.com/go-gourd/gourd"
	"gourd/app/event"
)

func main() {

	// 注册系统事件
	event.RegisterEvent()

	//database.GetDb()

	// 启动服务
	gourd.StartServer(true)

}

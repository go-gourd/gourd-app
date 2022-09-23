package event

import (
	"github.com/go-gourd/gourd"
	"gourd/app/http/router"
)

// InitEvent 初始化时执行
func InitEvent() {

	app := gourd.GetServer()

	//注册路由
	router.RegisterRoute(app)

}

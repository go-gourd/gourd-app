package event

import (
	"github.com/go-gourd/gourd"
	"gourd/app/cron"
	"gourd/app/http/middleware"
	"gourd/app/http/router"
)

// InitEvent 初始化时执行
func InitEvent() {

	app := gourd.GetServer()

	//注册Http全局中间件
	middleware.RegisterGlobalMiddleware(app)

	//注册Http路由
	router.RegisterRoute(app)

	//注册定时任务
	cron.RegisterCron()

}

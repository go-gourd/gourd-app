package initialize

import (
	"fmt"
	"gourd/internal/config"
	"gourd/internal/http/router"
	"log/slog"
	"net/http"
)

// InitHttpServer 启动http服务
func InitHttpServer() {

	// 初始化HTTP路由
	router.InitRouter()

	// 获取http配置文件
	conf, _ := config.GetHttpConfig()

	// 是否开启http服务
	if !conf.Enable {
		return
	}

	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	slog.Info("Started http server. " + addr)

	go func() {
		// 启动http服务
		r := router.GetRouter()
		err := http.ListenAndServe(addr, r)
		if err != nil {
			panic(err)
		}
	}()

}

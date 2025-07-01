package initialize

import (
	"app/internal/config"
	"app/internal/http/router"
	"fmt"
	"log/slog"
	"net/http"
)

// InitHttpServer 启动http服务
func InitHttpServer() {

	// 获取http配置文件
	conf, _ := config.GetHttpConfig()

	// 是否开启http服务
	if !conf.Enable {
		return
	}

	// 初始化HTTP路由
	router.InitRouter()

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

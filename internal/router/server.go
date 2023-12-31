package router

import (
	"fmt"
	"github.com/go-gourd/gourd/config"
	"github.com/go-gourd/gourd/log"
	"net/http"
)

// StartServer 启动http服务
func StartServer() {

	r := GetRouter()

	// 获取http配置
	conf := config.GetHttpConfig()

	// 是否开启http服务
	if !conf.Enable {
		return
	}

	addr := fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	log.Info("Started http server. " + addr)

	go func() {
		// 启动http服务
		err := http.ListenAndServe(addr, r)
		if err != nil {
			panic(err)
		}
	}()

}

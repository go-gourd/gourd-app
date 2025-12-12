package router

import (
	"app/internal/config"
	apiRouter "app/internal/http/api/router"
	"app/internal/http/common/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

// GetRouter 获取路由
// 若路由已注册，则直接返回，否则创建路由并返回
func GetRouter() *chi.Mux {
	if router != nil {
		return router
	}
	router = chi.NewRouter()
	return router
}

// InitRouter 注册路由
func InitRouter() {
	r := GetRouter()

	// 静态资源处理中间件
	conf, err := config.GetHttpConfig()
	if err == nil && conf.Static != "" {
		r.Use(middleware.StaticOrNext(conf.Static))
	}

	// 404响应
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 not found."))
	})

	// 主页
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello gourd!"))
	})

	// 注册api相关路由
	apiSubRouter := chi.NewRouter()
	apiRouter.Router(apiSubRouter)
	r.Mount("/api", apiSubRouter)

}

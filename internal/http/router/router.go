package router

import (
	"app/internal/config"
	apiRoute "app/internal/http/api/route"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
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

	// 404响应
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		// 若路由未定义，检测是否为静态资源
		conf, err := config.GetHttpConfig()
		// 若配置中有静态资源路径，尝试从该路径下查找资源
		if err == nil && conf.Static != "" {
			filepath := conf.Static + r.URL.Path
			//判断文件是否存在
			if info, err := os.Stat(filepath); err == nil && !info.IsDir() {
				http.ServeFile(w, r, filepath)
				return
			}
		}

		// 404响应内容
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 not found."))
	})

	// 主页
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello gourd!"))
	})
	// 注册api相关路由
	r.Mount("/api", chi.NewRouter().
		Group(apiRoute.RegisterRoute))

}

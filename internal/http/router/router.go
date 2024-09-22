package router

import (
	"github.com/go-chi/chi/v5"
	"gourd/internal/config"
	apiRoute "gourd/internal/http/api/route"
	"net/http"
	"os"
)

var router *chi.Mux

func GetRouter() *chi.Mux {
	if router != nil {
		return router
	}
	router = chi.NewRouter()
	return router
}

// Register 注册路由
func Register() {
	r := GetRouter()

	// 主页
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello gourd!"))
	})

	// 404响应
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {

		// 若路由未定义，检测是否为静态资源
		conf, err := config.GetHttpConfig()
		if err == nil && conf.Static != "" {
			filepath := conf.Static + r.URL.Path
			//判断文件是否存在
			_, err := os.Stat(filepath)
			if err == nil {
				http.ServeFile(w, r, filepath)
				return
			}
		}

		// 404响应内容
		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 not found."))
	})

	// 注册api相关路由
	r.Mount("/api", chi.NewRouter().
		Group(apiRoute.RegisterRoute))

}

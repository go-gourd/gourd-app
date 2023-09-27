package router

import (
	"github.com/go-chi/chi/v5"
	apiRoute "gourd/app/http/api/route"
	"net/http"
)

var router *chi.Mux

func GetRouter() *chi.Mux {

	if router != nil {
		return router
	}

	router = chi.NewRouter()

	return router
}

// RegisterRouter 初始化路由
func RegisterRouter() {

	r := GetRouter()

	//主页
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello gourd!"))
	})

	//404
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(404)
		_, _ = w.Write([]byte("404 not found."))
	})

	//注册api相关路由
	api := chi.NewRouter().Group(apiRoute.RegisterRoute)
	r.Mount("/api", api)

}

package router

import (
	"app/internal/http/api/controller"
	apiCtls "app/internal/http/api/controller"
	"github.com/go-chi/chi/v5"
)

// RegisterRouter 注册路由组
func RegisterRouter(r chi.Router) {

	r.Route("/user", func(r chi.Router) {
		c := apiCtls.User{}
		r.HandleFunc("/list", c.Info)
		r.HandleFunc("/add", c.Add)
	})

	testsCtl := controller.TestsCtl{}
	r.HandleFunc("/tests/test", testsCtl.Test)

}

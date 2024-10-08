package route

import (
	"github.com/go-chi/chi/v5"
	apiCtls "gourd/internal/http/api/ctl"
)

// RegisterRoute 注册路由组
func RegisterRoute(r chi.Router) {

	r.Route("/user", func(r chi.Router) {
		c := apiCtls.UserCtl{}
		r.HandleFunc("/list", c.Info)
		r.HandleFunc("/add", c.Add)
	})

	testsCtl := apiCtls.TestsCtl{}
	r.HandleFunc("/tests/test", testsCtl.Test)

}

package route

import (
	"github.com/go-chi/chi/v5"
	apiCtls "gourd/internal/app/api/ctl"
)

// RegisterRoute 注册路由组
func RegisterRoute(r chi.Router) {
	userCtl := apiCtls.UserController{}
	r.HandleFunc("/user/info", userCtl.Info)
	r.HandleFunc("/user/add", userCtl.Add)

	testsCtl := apiCtls.TestsController{}
	r.HandleFunc("/tests/test", testsCtl.Test)

}

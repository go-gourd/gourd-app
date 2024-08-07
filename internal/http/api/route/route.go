package route

import (
	"github.com/go-chi/chi/v5"
	apiCtls "gourd/internal/http/api/ctl"
)

// RegisterRoute 注册路由组
func RegisterRoute(r chi.Router) {
	userCtl := apiCtls.UserCtl{}
	r.HandleFunc("/user/info", userCtl.Info)
	r.HandleFunc("/user/add", userCtl.Add)

	testsCtl := apiCtls.TestsCtl{}
	r.HandleFunc("/tests/test", testsCtl.Test)

}

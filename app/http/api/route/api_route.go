package route

import (
	"github.com/go-chi/chi/v5"
	apiC "gourd/app/http/api/controller"
)

// RegisterRoute 注册路由组接口
func RegisterRoute(r chi.Router) {
	user := apiC.UserController{}
	r.HandleFunc("/user/info", user.Info)
	r.HandleFunc("/user/add", user.Add)

	test := apiC.TestController{}
	r.HandleFunc("/test/test", test.Test)

}

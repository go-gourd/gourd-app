package route

import (
	"github.com/go-chi/chi/v5"
	apiC "gourd/internal/app/api/controller"
)

// RegisterRoute 注册路由组接口
func RegisterRoute(r chi.Router) {
	user := apiC.UserController{}
	r.HandleFunc("/user/info", user.Info)
	r.HandleFunc("/user/add", user.Add)

	tests := apiC.TestsController{}
	r.HandleFunc("/tests/test", tests.Test)

}

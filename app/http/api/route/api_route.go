package route

import (
	"github.com/go-chi/chi/v5"
	apiC "gourd/app/http/api/controller"
)

// RegisterRoute 注册路由组接口
func RegisterRoute(r chi.Router) {
	h := apiC.UserController{}
	r.HandleFunc("/userInfo", h.UserInfo)
	r.HandleFunc("/userAdd", h.UserAdd)
}

package route

import (
	"github.com/gin-gonic/gin"
	apiC "gourd/app/http/api/controller"
)

// RegisterRoute 注册路由组接口
func RegisterRoute(r *gin.RouterGroup) {
	h := apiC.UserController{}
	r.GET("/userInfo", h.UserInfo)
	r.Any("/userAdd", h.UserAdd)
}

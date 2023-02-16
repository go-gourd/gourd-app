package api

import (
	"github.com/gin-gonic/gin"
	"gourd/app/http/api/controller"
)

// RegisterRoute 注册路由组接口
func RegisterRoute(r *gin.RouterGroup) {
	h := controller.HandlerUser{}
	r.GET("/users", h.GetUser)
}

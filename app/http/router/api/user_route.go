package api

import (
	"app/http/api/controller"
	"github.com/gin-gonic/gin"
)

// RegisterRoute 注册路由组接口
func RegisterRoute(r *gin.RouterGroup) {
	h := controller.HandlerUser{}
	r.GET("/users", h.GetUser)
}

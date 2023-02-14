package api

import (
	"app/http/api/user"
	"github.com/gin-gonic/gin"
)

// AddUserRouter 添加用户相关路由
func AddUserRouter(apiR *gin.RouterGroup) {
	apiR.Any("/user", user.Info)
}

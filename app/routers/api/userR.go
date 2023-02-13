package api

import (
	"app/http/api/user"
	"github.com/gin-gonic/gin"
)

func AddUserRouter(apiR *gin.RouterGroup) {
	apiR.Any("/user", user.Info)
}

package controller

import (
	"github.com/awesee/php2go/php"
	"github.com/gin-gonic/gin"
	"github.com/go-gourd/gourd/ghttp"
	"gourd/app/dal/query"
)

// UserController 用户控制器
type UserController struct {
}

// UserInfo 获取用户信息
func (*UserController) UserInfo(c *gin.Context) {

	user, _ := query.User.Where(query.User.ID.Eq(1)).First()
	if user != nil {
		user.Password = php.Md5(user.Password)
	}

	ghttp.Success(c, "", user)
}

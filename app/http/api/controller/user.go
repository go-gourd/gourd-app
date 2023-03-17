package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"gourd/app/dal/model"
	"gourd/app/dal/query"
	"gourd/app/http"
)

// UserController 用户控制器
type UserController struct {
	http.BaseController //继承基础控制器
}

// UserInfo 获取用户信息
func (ct *UserController) UserInfo(c *gin.Context) {

	//需要查询的字段
	fields := []field.Expr{
		query.User.ID,
		query.User.UserName,
	}

	user, _ := query.User.
		Where(query.User.ID.Eq(2)).
		Select(fields...).
		First()
	if user != nil {
	}

	//响应结果
	ct.Success(c, "", user)
}

// UserAdd 创建用户
func (ct *UserController) UserAdd(c *gin.Context) {

	user := model.User{
		UserName: "go_create",
	}

	err := query.User.Create(&user)
	if err != nil {
		fmt.Println("添加失败")
	}

	//响应结果
	ct.Success(c, "", user)
}

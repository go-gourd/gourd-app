package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gen/field"
	"gourd/app/http"
	"gourd/app/orm/model"
	"gourd/app/orm/query"
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
		Extend:   "{}",
	}

	err := query.User.Create(&user)
	if err != nil {
		fmt.Println("添加失败：" + err.Error())
	}

	//响应结果
	ct.Success(c, "", user)
}

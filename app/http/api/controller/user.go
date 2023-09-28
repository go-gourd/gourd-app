package controller

import (
	"fmt"
	"gorm.io/gen/field"
	"gourd/app/http/base"
	"gourd/app/orm/model"
	"gourd/app/orm/query"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	base.BaseController //继承基础控制器
}

// UserInfo 获取用户信息
func (ct *UserController) UserInfo(w http.ResponseWriter, _ *http.Request) {

	// 需要查询的字段
	fields := []field.Expr{
		query.User.ID,
		query.User.UserName,
	}

	user, _ := query.User.
		Where(query.User.ID.Eq(2)).
		Select(fields...).
		First()

	// 响应结果
	_ = ct.Success(w, "", user)
}

// UserAdd 创建用户
func (ct *UserController) UserAdd(w http.ResponseWriter, _ *http.Request) {

	user := model.User{
		UserName: "go_create",
	}

	err := query.User.Create(&user)
	if err != nil {
		fmt.Println("添加失败：" + err.Error())
	}

	// 响应结果
	_ = ct.Success(w, "", user)
}

package controller

import (
	"fmt"
	"gorm.io/gen/field"
	"gourd/internal/app"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	app.BaseController //继承基础控制器
}

// Info 获取用户信息
func (ctl *UserController) Info(w http.ResponseWriter, _ *http.Request) {

	// 需要查询的字段
	fields := []field.Expr{
		query.User.ID,
		query.User.Username,
	}

	user, _ := query.User.
		Where(query.User.ID.Eq(2)).
		Select(fields...).
		First()

	// 响应结果
	_ = ctl.Success(w, "", user)
}

// Add 创建用户
func (ctl *UserController) Add(w http.ResponseWriter, _ *http.Request) {

	user := model.User{
		Username: "go_create",
	}

	err := query.User.Create(&user)
	if err != nil {
		fmt.Println("添加失败：" + err.Error())
	}

	// 响应结果
	_ = ctl.Success(w, "", user)
}

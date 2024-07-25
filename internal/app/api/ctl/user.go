package controller

import (
	"gourd/internal/app"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"gourd/internal/repositories/user"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	app.BaseController //继承基础控制器
}

// Info 获取用户信息
func (ctl *UserController) Info(w http.ResponseWriter, r *http.Request) {

	repository := user.Repository{
		Ctx: r.Context(),
	}

	qu := query.User

	userData, _ := repository.Query().
		Where(
			qu.ID.Eq(1),
			qu.CreateTime.Eq(0),
		).
		Select(
			qu.ID,
			qu.UserName,
		).
		First()

	// 响应结果
	_ = ctl.Success(w, "", userData)
}

// Add 创建用户
func (ctl *UserController) Add(w http.ResponseWriter, r *http.Request) {

	repository := user.Repository{
		Ctx: r.Context(),
	}

	_ = repository.Begin()

	userData := model.User{
		UserName: "go_create",
	}

	err := repository.Create(&userData)
	if err != nil {
		_ = ctl.Fail(w, 1, "添加失败："+err.Error(), nil)
		_ = repository.Rollback()
		return
	}

	_ = repository.Commit()

	// 响应结果
	_ = ctl.Success(w, "", userData)
}

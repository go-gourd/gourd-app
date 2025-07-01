package controller

import (
	"app/internal/http/api/controller/common"
	"app/internal/orm/model"
	"app/internal/orm/query"
	"app/internal/repositories/user"
	"net/http"
)

// User 用户控制器
type User struct {
	common.Base //继承基础控制器
}

// Info 获取用户信息
func (ctl *User) Info(w http.ResponseWriter, r *http.Request) {

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
func (ctl *User) Add(w http.ResponseWriter, r *http.Request) {

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

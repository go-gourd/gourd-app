package controller

import (
	"gourd/internal/http/api/common"
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
	"net/http"
)

// UserCtl 用户控制器
type UserCtl struct {
	common.BaseCtl //继承基础控制器
}

// Info 获取用户信息
func (ctl *UserCtl) Info(w http.ResponseWriter, r *http.Request) {

	qu := query.User

	userData, _ := query.User.WithContext(r.Context()).
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
func (ctl *UserCtl) Add(w http.ResponseWriter, r *http.Request) {
	tx := query.Q.Begin()
	u := tx.User

	userData := model.User{
		UserName: "go_create",
	}

	// 创建用户
	err := u.Create(&userData)
	if err != nil {
		_ = ctl.Fail(w, 1, "添加失败："+err.Error(), nil)
		_ = tx.Rollback()
		return
	}

	// 其他操作...

	// 事务提交
	if err = tx.Commit(); err != nil {
		_ = ctl.Fail(w, 1, "添加失败："+err.Error(), nil)
	}

	// 响应结果
	_ = ctl.Success(w, "", userData)
}

package service

import (
	"gourd/internal/orm/model"
	"gourd/internal/orm/query"
)

type UserService struct {
}

// Add 添加用户
func (us *UserService) Add(user *model.User) (int32, error) {
	err := query.User.Create(user)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

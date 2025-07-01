package user

import (
	"app/internal/orm/model"
	"app/internal/orm/query"
	"app/internal/repositories/base"
	"context"
)

type Repository struct {
	base.Repository
	Ctx context.Context
}

func (r *Repository) Query() query.IUserDo {
	if r.Tx != nil {
		return r.Tx.User.WithContext(r.Ctx)
	}
	return query.User.WithContext(r.Ctx)
}

func (r *Repository) Create(ud *model.User) error {
	userQ := r.Query()
	// TODO: add more fields
	return userQ.Create(ud)
}

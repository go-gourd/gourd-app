package base

import (
	"app/internal/orm/query"
	"context"
	"errors"
)

// Repository is the base repository struct for all repositories.
type Repository struct {
	Tx  *query.QueryTx
	Ctx context.Context
}

// Begin starts a new transaction.
func (r *Repository) Begin() error {
	if r.Tx == nil {
		r.Tx = query.Q.Begin()
		return nil
	}
	return errors.New("transaction already started")
}

// Commit commits the current transaction.
func (r *Repository) Commit() error {
	if r.Tx == nil {
		return errors.New("no transaction started")
	}
	return r.Tx.Commit()
}

// Rollback rolls back the current transaction.
func (r *Repository) Rollback() error {
	if r.Tx == nil {
		return errors.New("no transaction started")
	}
	return r.Tx.Rollback()
}

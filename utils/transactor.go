package utils

import (
	"context"
	"gorm.io/gorm"
)

type Transactor interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type transactorImpl struct {
	db *gorm.DB
}

func NewTransactor(db *gorm.DB) Transactor {
	return &transactorImpl{db: db}
}

func (t *transactorImpl) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		ctxWithTx := InjectTx(ctx, tx)
		return fn(ctxWithTx)
	})
}

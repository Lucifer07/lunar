package utils

import (
	"context"
	"gorm.io/gorm"
)

type txKey struct{}

func InjectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func ExtractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

func GetQueryRunner(ctx context.Context, db *gorm.DB) *gorm.DB {
	tx := ExtractTx(ctx)
	if tx != nil {
		return tx
	}
	return db
}

package repositories

import (
	"context"
	"gorm.io/gorm"
)

type TransactionManagerRepository interface {
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

type transactionManagerRepository struct {
	db *gorm.DB
}

func NewTransactionManagerRepository(db *gorm.DB) TransactionManagerRepository {
	return &transactionManagerRepository{db}
}

func (r *transactionManagerRepository) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.WithContext(ctx).Transaction(fn)
}
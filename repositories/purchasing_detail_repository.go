package repositories

import (
	"context"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/gorm"
)

type PurchasingDetailRepository interface {
	Create(ctx context.Context, tx *gorm.DB, detail *models.PurchasingDetail) error
}

type purchasingDetailRepository struct {
	db *gorm.DB
}	

func NewPurchasingDetailRepository(db *gorm.DB) PurchasingDetailRepository {
	return &purchasingDetailRepository{db}
}

func (r *purchasingDetailRepository) Create(ctx context.Context, tx *gorm.DB, detail *models.PurchasingDetail) error {
	return tx.WithContext(ctx).Create(detail).Error
}
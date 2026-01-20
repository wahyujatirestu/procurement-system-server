package repositories

import (
	"context"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/gorm"
)

type PurchasingRepository interface {
	Create(ctx context.Context, tx *gorm.DB, purchasing *models.Purchasing) error
	FindAll() ([]models.Purchasing, error)
	UpdateGrandTotal(ctx context.Context, tx *gorm.DB, id uint, total float64) error
}

type purchasingRepository struct {
	db *gorm.DB
}

func NewPurchasingRepository(db *gorm.DB) PurchasingRepository {
	return &purchasingRepository{db}
}

func (r *purchasingRepository) Create(ctx context.Context, tx *gorm.DB, purchasing *models.Purchasing) error {
	return tx.WithContext(ctx).Create(purchasing).Error
}

func (r *purchasingRepository) FindAll() ([]models.Purchasing, error) {
	var purchasings []models.Purchasing
	
	err := r.db.Preload("Supplier").Preload("User").Preload("Details.Item").Order("date DESC").Find(&purchasings).Error
	return purchasings, err
}

func (r *purchasingRepository) UpdateGrandTotal(ctx context.Context, tx *gorm.DB, id uint, total float64) error {
	return tx.WithContext(ctx).Model(&models.Purchasing{}).Where("id = ?", id).Update("grand_total", total).Error
	
}
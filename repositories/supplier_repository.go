package repositories

import (
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *models.Supplier) error
	FindAll() ([]models.Supplier, error)
}

type supplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepository{db}
}

func (r *supplierRepository) Create(supplier *models.Supplier) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepository) FindAll() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

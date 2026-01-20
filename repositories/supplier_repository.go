package repositories

import (
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *models.Supplier) error
	FindAll() ([]models.Supplier, error)
	FindById(id uint) (*models.Supplier, error)
	Update(supplier *models.Supplier) error
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

func (r *supplierRepository) FindById(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.Where("id = ?", id).First(&supplier).Error
	return &supplier, err
}

func (r *supplierRepository) Update(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}
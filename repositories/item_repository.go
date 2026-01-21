package repositories

import (
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.Item) error
	FindAll() ([]models.Item, error)
	FindById(id uint) (*models.Item, error)
	Update(item *models.Item) error
	Delete(id uint) error
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *itemRepository) FindAll() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *itemRepository) FindById(id uint) (*models.Item, error) {
	var item models.Item
	err := r.db.Where("id = ?", id).First(&item).Error
	return &item, err
}

func (r *itemRepository) Update(item *models.Item) error {
	return r.db.Save(item).Error
}

func (r *itemRepository) Delete(id uint) error {
	return r.db.Delete(&models.Item{}, id).Error
}

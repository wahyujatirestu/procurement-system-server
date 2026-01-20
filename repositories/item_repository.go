package repositories

import (
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/gorm"
)

type ItemRepository interface {
	Create(item *models.Item) error
	FindAll() ([]models.Item, error)
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

package models

type Item struct {
	ID    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `gorm:"not null"`
	Stock int     `gorm:"not null"`
	Price float64 `gorm:"not null"`
}

func (Item) TableName() string {
	return "items"
}
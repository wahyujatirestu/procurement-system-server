package models

import (
	"gorm.io/gorm"
)

type Item struct {
	ID    		uint    		`gorm:"primaryKey;autoIncrement"`
	Name  		string  		`gorm:"not null"`
	Stock 		int     		`gorm:"not null"`
	Price 		float64 		`gorm:"not null"`
	DeletedAt 	gorm.DeletedAt 	`gorm:"index"`
}

func (Item) TableName() string {
	return "items"
}
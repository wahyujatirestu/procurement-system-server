package models

import "gorm.io/gorm"


type Supplier struct {
	ID      	uint   `gorm:"primaryKey;autoIncrement"`
	Name    	string `gorm:"not null"`
	Email   	string
	Address 	string
	Purchasings []Purchasing `gorm:"foreignKey:SupplierID"`
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
package models

import "time"

type Purchasing struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Date       time.Time
	SupplierID uint
	UserID     uint
	GrandTotal float64
	User       User
	Supplier   Supplier
	Details    []PurchasingDetail
}

func (Purchasing) TableName() string {
	return "purchasings"
}
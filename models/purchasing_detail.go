package models

type PurchasingDetail struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	PurchasingID uint
	ItemID       uint
	Qty          int
	SubTotal     float64
	Item         Item
}

func (PurchasingDetail) TableName() string {
	return "purchasing_details"
}

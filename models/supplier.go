package models

type Supplier struct {
	ID      	uint   `gorm:"primaryKey;autoIncrement"`
	Name    	string `gorm:"not null"`
	Email   	string
	Address 	string
	Purchasings []Purchasing `gorm:"foreignKey:SupplierID"`
}

func (Supplier) TableName() string {
	return "suppliers"
}
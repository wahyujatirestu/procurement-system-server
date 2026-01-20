package models

type User struct {
	ID 		 uint	`gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Purchasing []Purchasing	`gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "users"
}
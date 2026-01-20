package database

import (
	"fmt"
	"log"
	"github.com/wahyujatirestu/simple-procurement-system/config"
	"github.com/wahyujatirestu/simple-procurement-system/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg config.DBConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password,	cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Supplier{},
		&models.Item{},
		&models.Purchasing{},
		&models.PurchasingDetail{},
	)

	if err != nil {
		log.Fatal("failed to migrate database")
	}

	return db
}
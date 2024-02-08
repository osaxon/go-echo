package database

import (
	"log"

	"echo.osaxon/config"
	"echo.osaxon/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := config.Config("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})

	var count int64
	db.Model(&model.User{}).Count(&count)

	if count == 0 {
		db.Create(&model.User{Username: "user1", Email: "user1@example.com"})
		db.Create(&model.User{Username: "user1", Email: "user1@example.com"})
		db.Create(&model.User{Username: "user2", Email: "user2@example.com"})
		db.Create(&model.User{Username: "user3", Email: "user3@example.com"})
		db.Create(&model.User{Username: "user4", Email: "user4@example.com"})
	}
}

package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"spec-builder/backend/internal/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=specdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&models.Cabinet{}, &models.Equipment{})
}

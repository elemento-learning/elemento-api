package config

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.Quiz{})
	db.AutoMigrate(&models.Answer{})
	db.AutoMigrate(&models.Question{})

}

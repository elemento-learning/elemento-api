package config

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.UserResult{})
	db.AutoMigrate(&models.UserAnswer{})
	db.AutoMigrate(&models.FeedBack{})
}

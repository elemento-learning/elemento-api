package config

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.MagicCard{})
	db.AutoMigrate(&models.Bab{})
	db.AutoMigrate(&models.Senyawa{})
	db.AutoMigrate(&models.Modul{})
	db.AutoMigrate(&models.School{})
	db.AutoMigrate(&models.Class{})
}

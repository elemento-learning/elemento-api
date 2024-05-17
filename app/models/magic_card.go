package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MagicCard struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title    string    `gorm:"type:varchar(255);not null"`
	Subtitle string    `gorm:"type:text;not null"`
	Photo    string    `gorm:"type:text;not null"`
	PhotoURL string    `gorm:"type:text;not null"`
}

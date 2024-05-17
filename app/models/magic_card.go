package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MagicCard struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	NamaMolekul  string    `gorm:"type:varchar(255);not null"`
	UnsurMolekul string    `gorm:"type:varchar(255);not null"`
	Image        string    `gorm:"type:text;not null"`
	ImageUrl     string    `gorm:"type:text;not null"`
	Description  string    `gorm:"type:text;not null"`
	ListSenyawa  []Senyawa `gorm:"many2many:magic_card_senyawa;"`
}

type Senyawa struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Judul     string    `gorm:"type:varchar(255);not null"`
	Unsur     string    `gorm:"type:varchar(255);not null"`
	Deskripsi string    `gorm:"type:text;not null"`
}

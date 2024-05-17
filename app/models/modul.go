package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Modul struct {
	gorm.Model
	ModulID    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title      string    `gorm:"type:varchar(255);"`
	Subtitle   string    `gorm:"type:varchar(255);"`
	IsComplete bool      `gorm:"type:boolean;"`
	Image      string    `gorm:"type:varchar(255);"`
	ImageUrl   string    `gorm:"type:varchar(255);"`
	Babs       []Bab     `gorm:"foreignkey:TitleID"`
}

type Bab struct {
	gorm.Model
	TitleID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title         string    `gorm:"type:varchar(255);"`
	Description   string    `gorm:"type:text"`
	Task          string    `gorm:"type:text"`
	ResultStudent string    `gorm:"type:text"`
}

func (u *Modul) TableName() string {
	return "moduls"
}

// Path: app/models/modul.go

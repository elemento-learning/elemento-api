package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MagicCard struct {
	gorm.Model
	ID           uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	NamaMolekul  string    `gorm:"column:nama_molekul;type:varchar(255);"`
	UnsurMolekul string    `gorm:"column:unsur_molekul;type:varchar(255);"`
	Image        string    `gorm:"column:image;type:text;"`
	ImageUrl     string    `gorm:"column:image_url;type:text;"`
	Description  string    `gorm:"column:desc;type:text;"`
	ListSenyawa  []Senyawa `gorm:"foreignKey:MagicCardId;"`
}

type Senyawa struct {
	gorm.Model
	ID          uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	Judul       string    `gorm:"column:judul;type:varchar(255);"`
	Unsur       string    `gorm:"column:unsur;type:varchar(255);"`
	Deskripsi   string    `gorm:"column:desc;contype:text;"`
	MagicCardId uuid.UUID `gorm:"column:magic_card_senyawa;"`
}

func (u *MagicCard) TableName() string {
	return "magic_cards"
}

func (u *Senyawa) TableName() string {
	return "senyawas"
}

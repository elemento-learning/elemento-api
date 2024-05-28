package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbMagicCard struct {
	Conn *gorm.DB
}

func (db *dbMagicCard) CreateNewMagicCard(magicCard models.MagicCard) error {
	return db.Conn.Create(&magicCard).Error
}

func (db *dbMagicCard) GetMagicCardById(id uuid.UUID) (models.MagicCard, error) {
	var magicCard models.MagicCard
	err := db.Conn.Where("id = ?", id).First(&magicCard).Error
	return magicCard, err
}

func (db *dbMagicCard) GetAllMagicCard() ([]models.MagicCard, error) {
	var magicCards []models.MagicCard
	err := db.Conn.Find(&magicCards).Error
	return magicCards, err
}

func (db *dbMagicCard) UpdateMagicCard(magicCard models.MagicCard) error {

	err := db.Conn.Save(&magicCard).Error
	if err != nil {
		return err // Mengembalikan error yang terjadi saat menyimpan data
	}
	return nil
}

func (db *dbMagicCard) DeleteMagicCard(id uuid.UUID) error {
	var magicCard models.MagicCard
	err := db.Conn.Where("id = ?", id).Delete(&magicCard).Error
	return err
}

func (db *dbMagicCard) IntegrateSenyawaToMagicCard(magicCard models.MagicCard, senyawa models.Senyawa) error {
	err := db.Conn.Model(&magicCard).Association("ListSenyawa").Append(&senyawa)
	return err
}

func (db *dbMagicCard) RetrieveUpdatedMagicCardWithAssociatedSenyawa(uuid uuid.UUID) (models.MagicCard, error) {
	var magicCard models.MagicCard
	err := db.Conn.Preload("ListSenyawa").Where("id = ?", uuid).First(&magicCard).Error
	return magicCard, err
}

type MagicCardRepository interface {
	CreateNewMagicCard(magicCard models.MagicCard) error
	GetMagicCardById(id uuid.UUID) (models.MagicCard, error)
	GetAllMagicCard() ([]models.MagicCard, error)
	UpdateMagicCard(magicCard models.MagicCard) error
	IntegrateSenyawaToMagicCard(magicCard models.MagicCard, senyawa models.Senyawa) error
	RetrieveUpdatedMagicCardWithAssociatedSenyawa(uuid uuid.UUID) (models.MagicCard, error)
	DeleteMagicCard(id uuid.UUID) error
}

func NewMagicCardRepository(conn *gorm.DB) MagicCardRepository {
	return &dbMagicCard{Conn: conn}
}

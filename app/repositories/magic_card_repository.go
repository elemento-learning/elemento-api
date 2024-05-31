package repositories

import (
	"elemento-api/app/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbMagicCard struct {
	Conn *gorm.DB
}

func (db *dbMagicCard) CreateNewMagicCard(magicCard models.MagicCard) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error creating new MagicCard: %v", err)
		}
	}()
	err = db.Conn.Create(&magicCard).Error
	return err
}

func (db *dbMagicCard) GetMagicCardById(id uuid.UUID) (magicCard models.MagicCard, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting MagicCard by ID: %v", err)
		}
	}()
	err = db.Conn.Where("id = ?", id).First(&magicCard).Error
	return magicCard, err
}

func (db *dbMagicCard) GetAllMagicCard() (magicCards []models.MagicCard, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting all MagicCards: %v", err)
		}
	}()
	err = db.Conn.Find(&magicCards).Error
	return magicCards, err
}

func (db *dbMagicCard) UpdateMagicCard(magicCard models.MagicCard) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error updating MagicCard: %v", err)
		}
	}()
	err = db.Conn.Save(&magicCard).Error
	return err
}

func (db *dbMagicCard) DeleteMagicCard(id uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error deleting MagicCard: %v", err)
		}
	}()
	err = db.Conn.Where("id = ?", id).Delete(&models.MagicCard{}).Error
	return err
}

func (db *dbMagicCard) IntegrateSenyawaToMagicCard(magicCard models.MagicCard, senyawa models.Senyawa) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error integrating Senyawa to MagicCard: %v", err)
		}
	}()
	err = db.Conn.Model(&magicCard).Association("ListSenyawa").Append(&senyawa)
	return err
}

func (db *dbMagicCard) RetrieveUpdatedMagicCardWithAssociatedSenyawa(id uuid.UUID) (magicCard models.MagicCard, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error retrieving updated MagicCard with associated Senyawa: %v", err)
		}
	}()
	err = db.Conn.Preload("ListSenyawa").Where("id = ?", id).First(&magicCard).Error
	return magicCard, err
}

type MagicCardRepository interface {
	CreateNewMagicCard(magicCard models.MagicCard) error
	GetMagicCardById(id uuid.UUID) (models.MagicCard, error)
	GetAllMagicCard() ([]models.MagicCard, error)
	UpdateMagicCard(magicCard models.MagicCard) error
	IntegrateSenyawaToMagicCard(magicCard models.MagicCard, senyawa models.Senyawa) error
	RetrieveUpdatedMagicCardWithAssociatedSenyawa(id uuid.UUID) (models.MagicCard, error)
	DeleteMagicCard(id uuid.UUID) error
}

func NewMagicCardRepository(conn *gorm.DB) MagicCardRepository {
	return &dbMagicCard{Conn: conn}
}

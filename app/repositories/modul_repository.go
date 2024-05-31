package repositories

import (
	"elemento-api/app/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbModul struct {
	Conn *gorm.DB
}

func (db *dbModul) CreateNewModul(modul models.Modul) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error creating new Modul: %v", err)
		}
	}()
	err = db.Conn.Create(&modul).Error
	return err
}

func (db *dbModul) GetModulById(id uuid.UUID) (modul models.Modul, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting Modul by ID: %v", err)
		}
	}()
	err = db.Conn.Where("id = ?", id).First(&modul).Error
	return modul, err
}

func (db *dbModul) GetAllModul() (moduls []models.Modul, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting all Moduls: %v", err)
		}
	}()
	err = db.Conn.Find(&moduls).Error
	return moduls, err
}

func (db *dbModul) IntegrateBabToModul(modul models.Modul, bab models.Bab) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error integrating Bab to Modul: %v", err)
		}
	}()
	err = db.Conn.Model(&modul).Association("Babs").Append(&bab)
	return err
}

func (db *dbModul) RetrieveUpdatedModulWithAssociatedBab(id uuid.UUID) (modul models.Modul, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error retrieving updated Modul with associated Babs: %v", err)
		}
	}()
	err = db.Conn.Preload("Babs").First(&modul, id).Error
	return modul, err
}

func (db *dbModul) UpdateModul(modul models.Modul) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error updating Modul: %v", err)
		}
	}()
	err = db.Conn.Save(&modul).Error
	return err
}

func (db *dbModul) DeleteModul(id uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error deleting Modul: %v", err)
		}
	}()
	err = db.Conn.Where("id = ?", id).Delete(&models.Modul{}).Error
	return err
}

type ModulRepository interface {
	IntegrateBabToModul(modul models.Modul, bab models.Bab) error
	CreateNewModul(modul models.Modul) error
	GetModulById(id uuid.UUID) (models.Modul, error)
	GetAllModul() ([]models.Modul, error)
	UpdateModul(modul models.Modul) error
	DeleteModul(id uuid.UUID) error
	RetrieveUpdatedModulWithAssociatedBab(id uuid.UUID) (models.Modul, error)
}

func NewModulRepository(conn *gorm.DB) ModulRepository {
	return &dbModul{Conn: conn}
}

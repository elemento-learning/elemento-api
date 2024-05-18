package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbModul struct {
	Conn *gorm.DB
}

func (db *dbModul) CreateNewModul(modul models.Modul) error {
	return db.Conn.Create(&modul).Error
}

func (db *dbModul) GetModulById(id uuid.UUID) (models.Modul, error) {
	var modul models.Modul
	err := db.Conn.Where("id = ?", id).First(&modul).Error
	return modul, err
}

func (db *dbModul) GetAllModul() ([]models.Modul, error) {
	var moduls []models.Modul
	err := db.Conn.Find(&moduls).Error
	return moduls, err
}

func (db *dbModul) IntegrateBabToModul(modul models.Modul, bab models.Bab) error {
	err := db.Conn.Model(&modul).Association("Babs").Append(&bab)
	return err
}

func (db *dbModul) RetrieveUpdatedModulWithAssociatedBab(id uuid.UUID) (models.Modul, error) {
	var modul models.Modul
	err := db.Conn.Preload("Babs").Where("id = ?", id).First(&modul).Error
	return modul, err

}

func (db *dbModul) UpdateModul(modul models.Modul) error {

	err := db.Conn.Save(&modul).Error
	if err != nil {
		return err // Mengembalikan error yang terjadi saat menyimpan data
	}
	return nil
}

func (db *dbModul) DeleteModul(id uuid.UUID) error {
	var modul models.Modul
	err := db.Conn.Where("id = ?", id).Delete(&modul).Error
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

package repositories

import (
	"elemento-api/app/models"
	"log"

	"gorm.io/gorm"
)

// SenyawaRepository is a struct to define the repository for senyawa
type SenyawaRepository struct {
	DB *gorm.DB
}

// CreateNewSenyawa is a function to create new senyawa
func (r *SenyawaRepository) CreateNewSenyawa(senyawa models.Senyawa) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error creating new Senyawa: %v", err)
		}
	}()
	err = r.DB.Create(&senyawa).Error
	return err
}

// GetSenyawaById is a function to get senyawa by id
func (r *SenyawaRepository) GetSenyawaById(id uint) (senyawa models.Senyawa, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting Senyawa by ID: %v", err)
		}
	}()
	err = r.DB.Where("id = ?", id).First(&senyawa).Error
	return senyawa, err
}

// GetAllSenyawa is a function to get all senyawa
func (r *SenyawaRepository) GetAllSenyawa() (senyawa []models.Senyawa, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting all Senyawa: %v", err)
		}
	}()
	err = r.DB.Find(&senyawa).Error
	return senyawa, err
}

// UpdateSenyawa is a function to update senyawa
func (r *SenyawaRepository) UpdateSenyawa(senyawa models.Senyawa) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error updating Senyawa: %v", err)
		}
	}()
	err = r.DB.Save(&senyawa).Error
	return err
}

// DeleteSenyawa is a function to delete senyawa
func (r *SenyawaRepository) DeleteSenyawa(id uint) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error deleting Senyawa: %v", err)
		}
	}()
	err = r.DB.Where("id = ?", id).Delete(&models.Senyawa{}).Error
	return err
}

// GetSenyawaByMagicCardId is a function to get senyawa by magic card id
func (r *SenyawaRepository) GetSenyawaByMagicCardId(id uint) (senyawa []models.Senyawa, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting Senyawa by MagicCard ID: %v", err)
		}
	}()
	err = r.DB.Where("magic_card_id = ?", id).Find(&senyawa).Error
	return senyawa, err
}

// SenyawaRepositoryInterface is an interface for senyawa repository
type SenyawaRepositoryInterface interface {
	CreateNewSenyawa(senyawa models.Senyawa) error
	GetSenyawaById(id uint) (models.Senyawa, error)
	GetAllSenyawa() ([]models.Senyawa, error)
	UpdateSenyawa(senyawa models.Senyawa) error
	DeleteSenyawa(id uint) error
	GetSenyawaByMagicCardId(id uint) ([]models.Senyawa, error)
}

// NewSenyawaRepository is a function to create a new senyawa repository
func NewSenyawaRepository(db *gorm.DB) SenyawaRepository {
	return SenyawaRepository{
		DB: db,
	}
}

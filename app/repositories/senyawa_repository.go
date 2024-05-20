package repositories

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

// ModulRepository is a struct to define the repository for modul
type SenyawaRepository struct {
	DB *gorm.DB
}

// CreateNewSenyawa is a function to create new senyawa
func (r *SenyawaRepository) CreateNewSenyawa(senyawa models.Senyawa) error {
	return r.DB.Create(&senyawa).Error
}

// GetSenyawaById is a function to get senyawa by id
func (r *SenyawaRepository) GetSenyawaById(id uint) (models.Senyawa, error) {
	var senyawa models.Senyawa
	err := r.DB.Where("id = ?", id).First(&senyawa).Error
	return senyawa, err
}

// GetAllSenyawa is a function to get all senyawa
func (r *SenyawaRepository) GetAllSenyawa() ([]models.Senyawa, error) {
	var senyawa []models.Senyawa
	err := r.DB.Find(&senyawa).Error
	return senyawa, err
}

// UpdateSenyawa is a function to update senyawa
func (r *SenyawaRepository) UpdateSenyawa(senyawa models.Senyawa) error {
	err := r.DB.Save(&senyawa).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteSenyawa is a function to delete senyawa
func (r *SenyawaRepository) DeleteSenyawa(id uint) error {
	var senyawa models.Senyawa
	err := r.DB.Where("id = ?", id).Delete(&senyawa).Error
	return err
}

// GetSenyawaByMagicCardId is a function to get senyawa by magic card id
func (r *SenyawaRepository) GetSenyawaByMagicCardId(id uint) ([]models.Senyawa, error) {
	var senyawa []models.Senyawa
	err := r.DB.Where("magic_card_id = ?", id).Find(&senyawa).Error
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

func NewSenyawaRepository(db *gorm.DB) SenyawaRepository {
	return SenyawaRepository{
		DB: db,
	}
}

package repositories

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

// BabRepository is a struct to define the repository for bab
type BabRepository struct {
	DB *gorm.DB
}

// CreateNewBab is a function to create new bab
func (r *BabRepository) CreateNewBab(bab models.Bab) error {
	return r.DB.Create(&bab).Error
}

// GetBabById is a function to get bab by id
func (r *BabRepository) GetBabById(id uint) (models.Bab, error) {
	var bab models.Bab
	err := r.DB.Where("id = ?", id).First(&bab).Error
	return bab, err
}

// GetAllBab is a function to get all bab
func (r *BabRepository) GetAllBab() ([]models.Bab, error) {
	var babs []models.Bab
	err := r.DB.Find(&babs).Error
	return babs, err
}

// UpdateBab is a function to update bab
func (r *BabRepository) UpdateBab(bab models.Bab) error {
	err := r.DB.Save(&bab).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteBab is a function to delete bab
func (r *BabRepository) DeleteBab(id uint) error {
	var bab models.Bab
	err := r.DB.Where("id = ?", id).Delete(&bab).Error
	return err
}

func (r *BabRepository) GetBabByModulId(id uint) ([]models.Bab, error) {
	var babs []models.Bab
	err := r.DB.Where("modul_id = ?", id).Find(&babs).Error
	return babs, err
}

// BabRepositoryInterface is an interface for bab repository
type BabRepositoryInterface interface {
	CreateNewBab(bab models.Bab) error
	GetBabById(id uint) (models.Bab, error)
	GetAllBab() ([]models.Bab, error)
	UpdateBab(bab models.Bab) error
	DeleteBab(id uint) error
}

func NewBabRepository(db *gorm.DB) BabRepository {
	return BabRepository{
		DB: db,
	}
}

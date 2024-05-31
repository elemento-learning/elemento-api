package repositories

import (
	"elemento-api/app/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbSchool struct {
	Conn *gorm.DB
}

func (db *dbSchool) CreateNewSchool(school models.School) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error creating new School: %v", err)
		}
	}()
	err = db.Conn.Create(&school).Error
	return err
}

func (db *dbSchool) GetSchoolById(id uuid.UUID) (school models.School, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting School by ID: %v", err)
		}
	}()
	err = db.Conn.Where("id = ?", id).First(&school).Error
	return school, err
}

func (db *dbSchool) GetAllSchool() (schools []models.School, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting all Schools: %v", err)
		}
	}()
	err = db.Conn.Find(&schools).Error
	return schools, err
}

func (db *dbSchool) UpdateSchool(school models.School) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error updating School: %v", err)
		}
	}()
	err = db.Conn.Save(&school).Error
	return err
}

func (db *dbSchool) DeleteSchool(id uuid.UUID) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error deleting School: %v", err)
		}
	}()
	err = db.Conn.Where("id = ?", id).Delete(&models.School{}).Error
	return err
}

func (db *dbSchool) IntegrateClassToSchool(school models.School, class models.Class) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error integrating Class to School: %v", err)
		}
	}()
	err = db.Conn.Model(&school).Association("Classes").Append(&class)
	return err
}

type SchoolRepository interface {
	CreateNewSchool(school models.School) error
	GetSchoolById(id uuid.UUID) (models.School, error)
	GetAllSchool() ([]models.School, error)
	UpdateSchool(school models.School) error
	DeleteSchool(id uuid.UUID) error
	IntegrateClassToSchool(school models.School, class models.Class) error
}

// NewSchoolRepository is a function to create new school repository
func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &dbSchool{Conn: db}
}

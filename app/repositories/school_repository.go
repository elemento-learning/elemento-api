package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbSchool struct {
	Conn *gorm.DB
}

func (db *dbSchool) CreateNewSchool(school models.School) error {
	return db.Conn.Create(&school).Error
}

func (db *dbSchool) GetSchoolById(id uuid.UUID) (models.School, error) {
	var school models.School
	err := db.Conn.Where("id = ?", id).First(&school).Error
	return school, err
}

func (db *dbSchool) GetAllSchool() ([]models.School, error) {
	var schools []models.School
	err := db.Conn.Find(&schools).Error
	return schools, err
}

func (db *dbSchool) UpdateSchool(school models.School) error {

	err := db.Conn.Save(&school).Error
	if err != nil {
		return err // Mengembalikan error yang terjadi saat menyimpan data
	}
	return nil
}

func (db *dbSchool) DeleteSchool(id uuid.UUID) error {
	var school models.School
	err := db.Conn.Where("id = ?", id).Delete(&school).Error
	return err
}

func (db *dbSchool) IntegrateClassToSchool(school models.School, class models.Class) error {
	err := db.Conn.Model(&school).Association("Classes").Append(&class)
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

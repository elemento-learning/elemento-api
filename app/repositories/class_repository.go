package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbClass struct {
	Conn *gorm.DB
}

func (db *dbClass) CreateNewClass(class models.Class) error {
	return db.Conn.Create(&class).Error
}

func (db *dbClass) GetClassById(id uuid.UUID) (models.Class, error) {
	var class models.Class
	err := db.Conn.Where("id = ?", id).First(&class).Error
	return class, err
}

func (db *dbClass) GetClassBySchoolId(schoolId uuid.UUID) ([]models.Class, error) {
	var classes []models.Class
	err := db.Conn.Where("school_id = ?", schoolId).Find(&classes).Error
	return classes, err
}

func (db *dbClass) GetAllClass() ([]models.Class, error) {
	var classes []models.Class
	err := db.Conn.Find(&classes).Error
	return classes, err
}

func (db *dbClass) UpdateClass(class models.Class) error {

	err := db.Conn.Save(&class).Error
	if err != nil {
		return err // Mengembalikan error yang terjadi saat menyimpan data
	}
	return nil
}

func (db *dbClass) DeleteClass(id uuid.UUID) error {
	var class models.Class
	err := db.Conn.Where("id = ?", id).Delete(&class).Error
	return err
}

func NewClassRepository(conn *gorm.DB) ClassRepository {
	return &dbClass{Conn: conn}
}

type ClassRepository interface {
	CreateNewClass(class models.Class) error
	GetClassById(id uuid.UUID) (models.Class, error)
	GetAllClass() ([]models.Class, error)
	UpdateClass(class models.Class) error
	DeleteClass(id uuid.UUID) error
	GetClassBySchoolId(schoolId uuid.UUID) ([]models.Class, error)
}

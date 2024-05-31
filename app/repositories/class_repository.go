package repositories

import (
	"elemento-api/app/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbClass struct {
	Conn *gorm.DB
}

func (db *dbClass) CreateNewClass(class models.Class) error {
	err := db.Conn.Create(&class).Error
	if err != nil {
		log.Printf("Error creating new class: %v", err)
	}
	return err
}

func (db *dbClass) GetClassById(id uuid.UUID) (models.Class, error) {
	var class models.Class
	err := db.Conn.Where("id = ?", id).First(&class).Error
	if err != nil {
		log.Printf("Error getting class by ID: %v", err)
	}
	return class, err
}

func (db *dbClass) GetClassBySchoolId(schoolId uuid.UUID) ([]models.Class, error) {
	var classes []models.Class
	err := db.Conn.Where("school_id = ?", schoolId).Find(&classes).Error
	if err != nil {
		log.Printf("Error getting classes by school ID: %v", err)
	}
	return classes, err
}

func (db *dbClass) GetAllClass() ([]models.Class, error) {
	var classes []models.Class
	err := db.Conn.Find(&classes).Error
	if err != nil {
		log.Printf("Error getting all classes: %v", err)
	}
	return classes, err
}

func (db *dbClass) UpdateClass(class models.Class) error {
	err := db.Conn.Save(&class).Error
	if err != nil {
		log.Printf("Error updating class: %v", err)
	}
	return err
}

func (db *dbClass) DeleteClass(id uuid.UUID) error {
	err := db.Conn.Where("id = ?", id).Delete(&models.Class{}).Error
	if err != nil {
		log.Printf("Error deleting class: %v", err)
	}
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

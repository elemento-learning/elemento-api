package repositories

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

type UserResultRepository interface {
	CreateUserResult(result models.UserResult) error
	GetUserResult() ([]models.UserResult, error)
}

type userResultRepository struct {
	Conn *gorm.DB
}

func (db *userResultRepository) CreateUserResult(result models.UserResult) error {
	err := db.Conn.Create(&result).Error
	return err
}

func (db *userResultRepository) GetUserResult() ([]models.UserResult, error) {
	var results []models.UserResult
	err := db.Conn.Find(&results).Error
	return results, err
}

func NewUserResultRepository(db *gorm.DB) UserResultRepository {
	return &userResultRepository{Conn: db}
}

// Path: app/repositories/result_repository.go

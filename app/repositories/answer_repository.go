package repositories

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

// AnswerRepository is a struct to define the repository for answer
type dbAnswer struct {
	DB *gorm.DB
}

// CreateNewAnswer is a function to create new answer
func (r *dbAnswer) CreateNewAnswer(answer models.Answer) error {
	return r.DB.Create(&answer).Error
}

// GetAnswerById is a function to get answer by id
func (r *dbAnswer) GetAnswerById(id uint) (models.Answer, error) {
	var answer models.Answer
	err := r.DB.Where("id = ?", id).First(&answer).Error
	return answer, err
}

// GetAllAnswer is a function to get all answer
func (r *dbAnswer) GetAllAnswer() ([]models.Answer, error) {
	var answers []models.Answer
	err := r.DB.Find(&answers).Error
	return answers, err
}

// UpdateAnswer is a function to update answer
func (r *dbAnswer) UpdateAnswer(answer models.Answer) error {
	err := r.DB.Save(&answer).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteAnswer is a function to delete answer
func (r *dbAnswer) DeleteAnswer(id uint) error {
	var answer models.Answer
	err := r.DB.Where("id = ?", id).Delete(&answer).Error
	return err
}

// AnswerRepositoryInterface is an interface for answer repository
type AnswerRepository interface {
	CreateNewAnswer(answer models.Answer) error
	GetAnswerById(id uint) (models.Answer, error)
	GetAllAnswer() ([]models.Answer, error)
	UpdateAnswer(answer models.Answer) error
	DeleteAnswer(id uint) error
}

// NewAnswerRepository is a function to create new answer repository
func NewAnswerRepository(db *gorm.DB) AnswerRepository {
	return &dbAnswer{DB: db}
}

// Path: app/repositories/answer_repository.go

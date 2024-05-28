package repositories

import (
	"elemento-api/app/models"

	"gorm.io/gorm"
)

type quizRepository struct {
	Conn *gorm.DB
}

func (db *quizRepository) ListQuiz() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	//just select id, title, and status
	err := db.Conn.Select("id, title, status").Find(&quizzes).Error
	return quizzes, err
}

type QuizRepository interface {
	ListQuiz() ([]models.Quiz, error)
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &quizRepository{Conn: db}
}

package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type quizRepository struct {
	Conn *gorm.DB
}

func (db *quizRepository) CreateQuiz(quiz models.Quiz) error {
	err := db.Conn.Create(&quiz).Error
	return err
}

func (db *quizRepository) ListQuiz() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	//just select id, title, and status
	err := db.Conn.Select("id", "title", "status").Find(&quizzes).Error
	return quizzes, err
}

func (db *quizRepository) IntegrateQuestionWithQuiz(quiz models.Quiz, question models.Question) error {
	err := db.Conn.Model(&quiz).Association("Question").Append(&question)
	return err
}

func (db *quizRepository) RetrieveUpdatedQuizWithQuestionAndAnswer(quizID uuid.UUID) ([]models.Quiz, error) {
	var quizzes []models.Quiz
	err := db.Conn.Preload("Question.Answer").Find(&quizzes, quizID).Error
	return quizzes, err
}

// get quiz by id
func (db *quizRepository) GetQuizByID(quizID uuid.UUID) (models.Quiz, error) {
	var quiz models.Quiz
	err := db.Conn.Find(&quiz, quizID).Error
	return quiz, err
}

type QuizRepository interface {
	ListQuiz() ([]models.Quiz, error)
	IntegrateQuestionWithQuiz(quiz models.Quiz, question models.Question) error
	GetQuizByID(quizID uuid.UUID) (models.Quiz, error)
	RetrieveUpdatedQuizWithQuestionAndAnswer(quizID uuid.UUID) ([]models.Quiz, error)
	CreateQuiz(quiz models.Quiz) error
}

func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &quizRepository{Conn: db}
}

package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type questionRepository struct {
	Conn *gorm.DB
}

// Create question
func (db *questionRepository) CreateQuestion(question models.Question) error {
	err := db.Conn.Create(&question).Error
	return err
}

// select question by id
func (db *questionRepository) SelectQuestionByID(questionID uuid.UUID) (models.Question, error) {
	var question models.Question
	err := db.Conn.Find(&question, questionID).Error
	return question, err
}

// select question by quiz id
func (db *questionRepository) SelectQuestionByQuizID(quizID uuid.UUID) ([]models.Question, error) {
	var questions []models.Question
	err := db.Conn.Where("quiz_id = ?", quizID).Find(&questions).Error
	return questions, err
}

// update question
func (db *questionRepository) UpdateQuestion(question *models.Question) error {
	err := db.Conn.Save(question).Error
	return err
}

// delete question
func (db *questionRepository) DeleteQuestion(question *models.Question) error {
	err := db.Conn.Delete(question).Error
	return err
}

// Integrate answer with question
func (db *questionRepository) IntegrateAnswerWithQuestion(question models.Question, answer models.Answer) error {
	err := db.Conn.Model(&question).Association("Answer").Append(&answer)
	return err
}

type QuestionRepository interface {
	CreateQuestion(question models.Question) error
	SelectQuestionByID(questionID uuid.UUID) (models.Question, error)
	SelectQuestionByQuizID(quizID uuid.UUID) ([]models.Question, error)
	IntegrateAnswerWithQuestion(question models.Question, answer models.Answer) error
	UpdateQuestion(question *models.Question) error
	DeleteQuestion(question *models.Question) error
}

func NewQuestionRepository(db *gorm.DB) QuestionRepository {
	return &questionRepository{Conn: db}
}

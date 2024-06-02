package repositories

import (
	"elemento-api/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BabRepository is a struct to define the repository for bab
type feedbackRepository struct {
	DB *gorm.DB
}

// CreateNewFeedback is a function to create new feedback
func (r *feedbackRepository) CreateNewFeedback(feedback models.FeedBack) error {
	return r.DB.Create(&feedback).Error
}

// GetFeedbackById is a function to get feedback by id
func (r *feedbackRepository) GetFeedbackById(id uuid.UUID) (models.FeedBack, error) {
	var feedback models.FeedBack
	err := r.DB.Where("id = ?", id).First(&feedback).Error
	return feedback, err
}

func (r *feedbackRepository) GetFeedbackByStudentId(id uuid.UUID) ([]models.FeedBack, error) {
	var feedbacks []models.FeedBack
	err := r.DB.Where("student_id = ?", id).Find(&feedbacks).Error
	return feedbacks, err
}

// GetAllFeedback is a function to get all feedback
func (r *feedbackRepository) GetAllFeedback() ([]models.FeedBack, error) {
	var feedbacks []models.FeedBack
	err := r.DB.Find(&feedbacks).Error
	return feedbacks, err
}

// UpdateFeedback is a function to update feedback
func (r *feedbackRepository) UpdateFeedback(feedback models.FeedBack) error {
	return r.DB.Save(&feedback).Error
}

// DeleteFeedback is a function to delete feedback
func (r *feedbackRepository) DeleteFeedback(id uuid.UUID) error {
	var feedback models.FeedBack
	err := r.DB.Where("id = ?", id).Delete(&feedback).Error
	return err
}

type FeedbackRepository interface {
	CreateNewFeedback(feedback models.FeedBack) error
	GetFeedbackById(id uuid.UUID) (models.FeedBack, error)
	GetAllFeedback() ([]models.FeedBack, error)
	UpdateFeedback(feedback models.FeedBack) error
	DeleteFeedback(id uuid.UUID) error
	GetFeedbackByStudentId(id uuid.UUID) ([]models.FeedBack, error)
}

func NewFeedbackRepository(db *gorm.DB) FeedbackRepository {
	return &feedbackRepository{
		DB: db,
	}
}

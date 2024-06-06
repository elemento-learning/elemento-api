package services

import (
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// FeedbackService is a struct to define the service for feedback
type FeedbackService struct {
	feedbackRepository repositories.FeedbackRepository
	userRepository     repositories.UserRepository
}

func (service *FeedbackService) DeleteFeedback(id uuid.UUID, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	err := service.feedbackRepository.DeleteFeedback(id)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal menghapus feedback",
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil menghapus feedback",
		Data:       nil,
	}
}

func (service *FeedbackService) CreateFeedback(teacherId uuid.UUID, feedBack string, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}
	userId, err := utils.ParseDataId(bearerToken)
	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	user, err := service.userRepository.GetUserById(userId)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mendapatkan data user",
			Data:       nil,
		}
	}

	feedback := models.FeedBack{
		StudentID:  user.IdUser,
		TeacherID:  teacherId,
		FeedBack:   feedBack,
		FeedBackID: uuid.New(),
		Category:   "Question",
	}

	err = service.feedbackRepository.CreateNewFeedback(feedback)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal membuat feedback",
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil membuat feedback",
		Data:       nil,
	}
}

func (service *FeedbackService) GetFeedbacks(bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	feedbacks, err := service.feedbackRepository.GetAllFeedback()
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mendapatkan feedback",
			Data:       nil,
		}
	}

	feedbacksData := []map[string]interface{}{}
	for _, feedback := range feedbacks {
		student, err := service.userRepository.GetUserById(feedback.StudentID)
		if err != nil {
			return utils.Response{
				StatusCode: 500,
				Messages:   "Gagal mendapatkan data student",
				Data:       nil,
			}

		}
		teacher, err := service.userRepository.GetUserById(feedback.TeacherID)
		if err != nil {
			return utils.Response{
				StatusCode: 500,
				Messages:   "Gagal mendapatkan data teacher",
				Data:       nil,
			}
		}

		feedbackData := map[string]interface{}{
			"feedback_id": feedback.FeedBackID,
			"student":     student.Fullname,
			"teacher":     teacher.Fullname,
			"feedback":    feedback.FeedBack,
			"category":    feedback.Category,
		}

		feedbacksData = append(feedbacksData, feedbackData)
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil mendapatkan feedback",
		Data:       feedbacksData,
	}
}

func (service *FeedbackService) GetFeedbacksByStudentId(studentId uuid.UUID, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	feedbacks, err := service.feedbackRepository.GetFeedbackByStudentId(studentId)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mendapatkan feedback",
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil mendapatkan feedback",
		Data:       feedbacks,
	}
}

func NewFeedbackService(db *gorm.DB) FeedbackService {
	return FeedbackService{
		feedbackRepository: repositories.NewFeedbackRepository(db),
		userRepository:     repositories.NewDBUserRepository(db),
	}
}

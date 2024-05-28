package services

import (
	"elemento-api/app/repositories"
	"elemento-api/utils"

	"gorm.io/gorm"
)

type QuizService struct {
	quizRepository repositories.QuizRepository
}

func (service *QuizService) ListQuiz(bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	quizzes, err := service.quizRepository.ListQuiz()
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mendapatkan data quiz",
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil mendapatkan data quiz",
		Data:       quizzes,
	}

}

func NewQuizService(db *gorm.DB) QuizService {
	repo := repositories.NewQuizRepository(db)
	return QuizService{
		quizRepository: repo,
	}
}

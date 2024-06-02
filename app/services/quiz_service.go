package services

import (
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"
	"sort"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuizService struct {
	quizRepository       repositories.QuizRepository
	questionRepository   repositories.QuestionRepository
	answerRepository     repositories.AnswerRepository
	userResultRepository repositories.UserResultRepository
	userRepository       repositories.UserRepository
}

// Get Leaderboard
func (service *QuizService) GetLeaderboard(bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var leaderboard []utils.StudentScore
	users, err := service.userResultRepository.GetUserResult()
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mendapatkan data users",
			Data:       nil,
		}
	}

	for _, user := range users {
		userName, err := service.userRepository.GetUserById(user.UserID)
		if err != nil {
			return utils.Response{
				StatusCode: 500,
				Messages:   "Gagal mendapatkan data user",
				Data:       nil,
			}
		}

		leaderboard = append(leaderboard, utils.StudentScore{
			Score:     user.Score,
			StudentID: user.UserID,
			Name:      userName.Fullname,
		})
	}

	sort.SliceStable(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil mendapatkan data leaderboard",
		Data:       leaderboard,
	}
}

// Submit Quiz
func (service *QuizService) SubmitQuiz(quizID uuid.UUID, answers []utils.UserAnswerRequest, bearerToken string) utils.Response {
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

	var response utils.Response
	if quizID == uuid.Nil || len(answers) == 0 {
		response.StatusCode = 400
		response.Messages = "QuizID, UserID, dan Answer tidak boleh kosong"
		response.Data = nil
		return response
	}

	quiz, err := service.quizRepository.RetrieveUpdatedQuizWithQuestionAndAnswer(quizID)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan data quiz"
		response.Data = nil
		return response
	}

	var score int
	var newUserAnswer []models.UserAnswer
	var newUserResult models.UserResult = models.UserResult{
		UserResultID: uuid.New(),
		UserID:       userId,
		QuizID:       quizID,
		CountAnswer:  len(answers),
	}

	for _, answer := range answers {
		for _, question := range quiz.Question {
			if question.QuestionID == answer.QuestionID {
				for _, correctAnswer := range question.Answer {
					if correctAnswer.AnswerID == answer.AnswerID {
						score += 1
						newUserAnswer = append(newUserAnswer, models.UserAnswer{
							UserAnswerID:       uuid.New(),
							TitleQuestion:      question.Question,
							UserAnswerTitle:    correctAnswer.AnswerTitle,
							UserAnswerSubtitle: correctAnswer.AnswerSubtitle,
						})
					}
				}
			}
		}
	}
	newUserResult.Score = utils.CalculateScore(score, len(answers))
	newUserResult.Answer = newUserAnswer
	err = service.userResultRepository.CreateUserResult(newUserResult)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal submit quiz"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil submit quiz"
	response.Data = newUserResult
	return response
}

// Create Answer And Integrate To Question
func (service *QuizService) CreateAnswerAndIntegrateToQuestion(questionID uuid.UUID, answer utils.AnswerRequest, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	if answer.AnswerTitle == "" || questionID == uuid.Nil {
		response.StatusCode = 400
		response.Messages = "Answer dan question_id tidak boleh kosong"
		response.Data = nil
		return response
	}

	newAnswer := models.Answer{
		AnswerTitle:    answer.AnswerTitle,
		AnswerSubtitle: answer.AnswerSubtitle,
		QuestionID:     questionID,
		AnswerID:       uuid.New(),
	}

	err := service.answerRepository.CreateNewAnswer(newAnswer)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat answer " + err.Error()
		response.Data = nil
		return response
	}

	question, err := service.questionRepository.SelectQuestionByID(questionID)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan question"
		response.Data = nil
		return response
	}

	service.questionRepository.IntegrateAnswerWithQuestion(question, newAnswer)

	response.StatusCode = 200
	response.Messages = "Berhasil membuat answer"
	response.Data = newAnswer
	return response
}

// Create Question And Integrate To Quiz
func (service *QuizService) CreateQuestionAndIntegrateToQuiz(quizID uuid.UUID, question utils.QuestionRequest, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	if question.Question == "" || quizID == uuid.Nil {
		response.StatusCode = 400
		response.Messages = "Question dan quiz_id tidak boleh kosong"
		response.Data = nil
		return response
	}

	newQuestion := models.Question{
		Question:   question.Question,
		QuizID:     quizID,
		QuestionID: uuid.New(),
	}

	err := service.questionRepository.CreateQuestion(newQuestion)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat question"
		response.Data = nil
		return response
	}

	quiz, err := service.quizRepository.GetQuizByID(quizID)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan quiz"
		response.Data = nil
		return response
	}
	service.quizRepository.IntegrateQuestionWithQuiz(quiz, newQuestion)

	response.StatusCode = 200
	response.Messages = "Berhasil membuat question"
	response.Data = newQuestion
	return response
}

// Create quiz
func (service *QuizService) CreateQuiz(quiz utils.QuizRequest, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}
	var response utils.Response
	if quiz.Title == "" || quiz.Status == "" {
		response.StatusCode = 400
		response.Messages = "Title dan status tidak boleh kosong"
		response.Data = nil
		return response
	}

	newQuiz := models.Quiz{
		Title:  quiz.Title,
		Status: quiz.Status,
	}

	newQuiz.QuizID = uuid.New()
	err := service.quizRepository.CreateQuiz(newQuiz)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat quiz"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil membuat quiz"
	response.Data = newQuiz
	return response
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

func (service *QuizService) GetQuestionQuiz(quizID uuid.UUID, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	questions, err := service.quizRepository.RetrieveUpdatedQuizWithQuestionAndAnswer(quizID)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mendapatkan data question",
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil mendapatkan data question",
		Data:       questions,
	}
}

func NewQuizService(db *gorm.DB) QuizService {
	return QuizService{
		quizRepository:       repositories.NewQuizRepository(db),
		questionRepository:   repositories.NewQuestionRepository(db),
		answerRepository:     repositories.NewAnswerRepository(db),
		userResultRepository: repositories.NewUserResultRepository(db),
		userRepository:       repositories.NewDBUserRepository(db),
	}
}

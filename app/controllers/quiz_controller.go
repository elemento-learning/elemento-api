package controllers

import (
	"elemento-api/app/services"
	"elemento-api/utils"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type QuizController struct {
	quizService services.QuizService
	vl          vl.Validate
}

func NewQuizController(db *gorm.DB) QuizController {
	service := services.NewQuizService(db)
	controller := QuizController{
		quizService: service,
		vl:          *vl.New(),
	}
	return controller
}

func (controller *QuizController) GetQuestionQuiz(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	quizID := c.Param("id")
	uuidQuiz := uuid.MustParse(quizID)
	response := controller.quizService.GetQuestionQuiz(uuidQuiz, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *QuizController) ListQuiz(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.quizService.ListQuiz(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *QuizController) CreateQuiz(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		Title  string `json:"title" validate:"required"`
		Status string `json:"status" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	quizRequest := utils.QuizRequest{
		Title:  payloadValidator.Title,
		Status: payloadValidator.Status,
	}

	response := controller.quizService.CreateQuiz(quizRequest, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *QuizController) CreateQuestion(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	quizid := c.Param("id")

	type payload struct {
		Question string `json:"question" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	questionRequest := utils.QuestionRequest{
		Question: payloadValidator.Question,
	}

	quizUuid := uuid.MustParse(quizid)

	response := controller.quizService.CreateQuestionAndIntegrateToQuiz(quizUuid, questionRequest, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *QuizController) CreateAnswer(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	questionId := c.Param("id")

	type payload struct {
		AnswerTitle    string `json:"answer_title" validate:"required"`
		AnswerSubtitle string `json:"answer_subtitle" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	answerRequest := utils.AnswerRequest{
		AnswerTitle:    payloadValidator.AnswerTitle,
		AnswerSubtitle: payloadValidator.AnswerSubtitle,
	}
	uuidQuestion := uuid.MustParse(questionId)
	response := controller.quizService.CreateAnswerAndIntegrateToQuestion(uuidQuestion, answerRequest, token)
	return c.JSON(response.StatusCode, response)
}

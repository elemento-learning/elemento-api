package controllers

import (
	"elemento-api/app/services"
	"strings"

	vl "github.com/go-playground/validator/v10"
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

func (controller *QuizController) ListQuiz(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.quizService.ListQuiz(token)
	return c.JSON(response.StatusCode, response)
}

package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"elemento-api/app/controllers"
)

func QuizRoutes(apiv1 *echo.Group, db *gorm.DB) {
	quizController := controllers.NewQuizController(db)

	apiv1.GET("/quiz", quizController.ListQuiz)

}

package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"elemento-api/app/controllers"
)

func QuizRoutes(apiv1 *echo.Group, db *gorm.DB) {
	quizController := controllers.NewQuizController(db)

	apiv1.GET("/quiz", quizController.ListQuiz)
	apiv1.GET("/question-quiz/:id", quizController.GetQuestionQuiz)
	apiv1.POST("/quiz", quizController.CreateQuiz)
	apiv1.POST("/quiz/:id/question", quizController.CreateQuestion)
	apiv1.POST("/quiz/question/:id/answer", quizController.CreateAnswer)
	apiv1.GET("/question-quiz/:id", quizController.GetQuestionQuiz)
	apiv1.POST("/quiz/submit", quizController.SubmitQuiz)
	apiv1.GET("/quiz/leaderboard", quizController.GetLeaderboard)

}

package routes

import (
	"elemento-api/app/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// RouteFeedback is a function to define the routes for feedback
func RouteFeedback(apiv1 *echo.Group, db *gorm.DB) {
	feedbackController := controllers.NewFeedbackController(db)

	apiv1.POST("/feedback", feedbackController.CreateFeedback)
	apiv1.GET("/feedback", feedbackController.GetFeedbacks)
	apiv1.GET("/feedback/:id", feedbackController.GetFeedbackById)
	apiv1.DELETE("/feedback/:id", feedbackController.DeleteFeedback)
}

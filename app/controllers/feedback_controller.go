package controllers

import (
	"strings"

	"elemento-api/app/services"

	vl "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// FeedbackController is a struct to define the feedback controller
type FeedbackController struct {
	feedbackService services.FeedbackService
	vl              vl.Validate
}

func NewFeedbackController(db *gorm.DB) *FeedbackController {
	service := services.NewFeedbackService(db)
	controller := FeedbackController{
		feedbackService: service,
		vl:              *vl.New(),
	}
	return &controller
}

func (controller *FeedbackController) CreateFeedback(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		Feedback  string    `json:"feedback" validate:"required"`
		StudentID uuid.UUID `json:"student_id" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	feedback := payloadValidator.Feedback
	studentId := payloadValidator.StudentID

	response := controller.feedbackService.CreateFeedback(studentId, feedback, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *FeedbackController) GetFeedbacks(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.feedbackService.GetFeedbacks(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *FeedbackController) GetFeedbackById(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid := uuid.MustParse(id)

	response := controller.feedbackService.GetFeedbacksByStudentId(uuid, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *FeedbackController) DeleteFeedback(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid := uuid.MustParse(id)

	response := controller.feedbackService.DeleteFeedback(uuid, token)
	return c.JSON(response.StatusCode, response)
}

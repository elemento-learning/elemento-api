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

type SchoolController struct {
	schoolService services.SchoolService
	validate      vl.Validate
}

func NewSchoolController(db *gorm.DB) SchoolController {
	service := services.NewSchoolService(db)
	controller := SchoolController{
		schoolService: service,
		validate:      *vl.New(),
	}
	return controller
}

func (controller *SchoolController) CreateNewSchool(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		Name     string `form:"name" validate:"required"`
		Location string `form:"location" validate:"required"`
	}

	payloadValidator := new(payload)
	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	schoolRequest := utils.SchoolRequest{
		Name:     payloadValidator.Name,
		Location: payloadValidator.Location,
	}

	response := controller.schoolService.CreateNewSchool(schoolRequest, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *SchoolController) GetSchoolById(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	schoolId := c.Param("id")

	uuid, err := uuid.Parse(schoolId)
	if err != nil {
		return c.JSON(400, "Invalid school id")
	}

	response := controller.schoolService.GetSchoolById(uuid, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *SchoolController) GetAllSchools(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.schoolService.GetAllSchool(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *SchoolController) IntegrateClassToSchool(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		SchoolId uuid.UUID `form:"schoolId" validate:"required"`
		Nama     string    `form:"name" validate:"required"`
		Location string    `form:"location" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	classRequest := utils.ClassRequest{
		Name:     payloadValidator.Nama,
		Location: payloadValidator.Location,
	}

	response := controller.schoolService.IntegrateClassToSchool(payloadValidator.SchoolId, classRequest, token)
	return c.JSON(response.StatusCode, response)
}

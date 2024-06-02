package controllers

import (
	"elemento-api/app/services"
	"elemento-api/utils"
	"net/http"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ModulController struct {
	modulService services.ModulService
	validate     vl.Validate
}

func NewModulController(db *gorm.DB) ModulController {
	service := services.NewModulService(db)
	controller := ModulController{
		modulService: service,
		validate:     *vl.New(),
	}
	return controller
}

func (controller *ModulController) DeleteBab(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uintId, err := utils.StringToUint(id)

	if err != nil {
		return c.JSON(400, "Invalid ID")
	}

	response := controller.modulService.DeleteBab(uintId, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *ModulController) DeleteModul(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, "Invalid UUID")
	}

	response := controller.modulService.DeleteModul(uuid, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *ModulController) CreateNewModul(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		Title       string `form:"titleModul" validate:"required"`
		IsComplete  bool   `form:"isComplete" validate:"required"`
		YoutubeLink string `form:"youtubeLink" validate:"required"`
		Subtitle    string `form:"subtitleModul" validate:"required"`
	}

	if err := c.Request().ParseMultipartForm(1024); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	alias := c.Request().FormValue("alias")
	uploadedFile, handler, err := c.Request().FormFile("file")
	if err != nil {
		return c.JSON(400, err.Error())
	}

	photoRequest := utils.UploadedPhoto{
		Alias:   alias,
		File:    uploadedFile,
		Handler: handler,
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	ModulRequest := utils.ModulRequest{
		Title:      payloadValidator.Title,
		IsComplete: payloadValidator.IsComplete,
		Subtitle:   payloadValidator.Subtitle,
	}

	if err := controller.validate.Struct(ModulRequest); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.modulService.CreateNewModul(ModulRequest, token, photoRequest)
	return c.JSON(response.StatusCode, response)
}

func (controller *ModulController) GetModul(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.modulService.GetModul(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *ModulController) CreateBabAndIntegrateToModul(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type payload struct {
		Title       string    `form:"titleBab" validate:"required"`
		Description string    `form:"descriptionBab" validate:"required"`
		Task        string    `form:"taskBab" validate:"required"`
		ModulId     uuid.UUID `form:"modulId" validate:"required"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	BabRequest := utils.BabRequest{
		Title:       payloadValidator.Title,
		Description: payloadValidator.Description,
		Task:        payloadValidator.Task,
	}

	if err := controller.validate.Struct(BabRequest); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.modulService.CreateBabAndIntegrateToModul(payloadValidator.ModulId, token, BabRequest)
	return c.JSON(response.StatusCode, response)
}

func (controller *ModulController) GetModulById(c echo.Context) error {
	id := c.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(400, "Invalid UUID")
	}

	response := controller.modulService.GetModulById(uuid)
	return c.JSON(response.StatusCode, response)
}

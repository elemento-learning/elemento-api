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

func (controller *ModulController) CreateNewModul(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type Bab struct {
		Title         string `form:"title" validate:"required"`
		Description   string `form:"description" validate:"required"`
		Task          string `form:"task"`
		ResultStudent string `form:"resultStudent"`
	}

	type payload struct {
		Title      string `form:"title" validate:"required"`
		IsComplete bool   `form:"isComplete" validate:"required"`
		Subtitle   string `form:"subtitle" validate:"required"`
		Image      string `form:"image" validate:"required"`
		Babs       []Bab  `form:"babs" validate:"required"`
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

	var babs []utils.BabRequest

	for _, bab := range payloadValidator.Babs {
		bab := utils.BabRequest{
			Title:         bab.Title,
			Description:   bab.Description,
			Task:          bab.Task,
			ResultStudent: bab.ResultStudent,
		}

		if err := controller.validate.Struct(bab); err != nil {
			return c.JSON(400, err.Error())
		}

		babs = append(babs, bab)
	}

	ModulRequest := utils.ModulRequest{
		Title:      payloadValidator.Title,
		IsComplete: payloadValidator.IsComplete,
		Subtitle:   payloadValidator.Subtitle,
		Image:      payloadValidator.Image,
		Babs:       babs,
	}

	if err := controller.validate.Struct(ModulRequest); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.modulService.CreateNewModul(ModulRequest, token, photoRequest)
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

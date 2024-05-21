package controllers

import (
	"elemento-api/app/services"
	"elemento-api/utils"
	"net/http"
	"strings"

	vl "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MagicCardController struct {
	magicCardService services.MagicCardService
	validate         vl.Validate
}

func NewMagicCardController(db *gorm.DB) MagicCardController {
	service := services.NewMagicCardService(db)
	controller := MagicCardController{
		magicCardService: service,
		validate:         *vl.New(),
	}
	return controller
}

func (controller *MagicCardController) CreateMagicCard(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type Senyawa struct {
		Judul     string `form:"judul" validate:"required"`
		Unsur     string `form:"unsur" validate:"required"`
		Deskripsi string `form:"deskripsi" validate:"required"`
	}

	type payload struct {
		NamaMolekul  string `form:"namaMolekul" validate:"required"`
		UnsurMolekul string `form:"unsurMolekul" validate:"required"`
		Image        string `form:"image" validate:"required"`
		ImageUrl     string `form:"imageUrl" validate:"required"`
		Description  string `form:"description" validate:"required"`
		ListSenyawa  []Senyawa
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

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	var listSenyawa []utils.SenyawaRequest
	for _, senyawa := range payloadValidator.ListSenyawa {
		senyawa := utils.SenyawaRequest{
			Judul:     senyawa.Judul,
			Unsur:     senyawa.Unsur,
			Deskripsi: senyawa.Deskripsi,
		}

		if err := controller.validate.Struct(senyawa); err != nil {
			return c.JSON(400, err.Error())
		}

		listSenyawa = append(listSenyawa, senyawa)
	}

	magicCardReq := utils.MagicCardRequest{
		NamaMolekul:  payloadValidator.NamaMolekul,
		UnsurMolekul: payloadValidator.UnsurMolekul,
		Image:        payloadValidator.Image,
		ImageUrl:     payloadValidator.ImageUrl,
		Description:  payloadValidator.Description,
		ListSenyawa:  listSenyawa,
	}

	response := controller.magicCardService.CreateMagicCard(magicCardReq, token, photoRequest)

	return c.JSON(response.StatusCode, response)
}

func (controller *MagicCardController) GetMagicCardById(c echo.Context) error {
	id := c.Param("id")
	uuid, err := utils.ParseDataId(id)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.magicCardService.GetMagicCardById(uuid, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *MagicCardController) GetAllMagicCard(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.magicCardService.GetAllMagicCard(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *MagicCardController) UpdateMagicCard(c echo.Context) error {
	id := c.Param("id")
	uuid, err := utils.ParseDataId(id)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	type Senyawa struct {
		Judul     string `form:"judul" validate:"required"`
		Unsur     string `form:"unsur" validate:"required"`
		Deskripsi string `form:"deskripsi" validate:"required"`
	}

	type payload struct {
		NamaMolekul  string `form:"namaMolekul" validate:"required"`
		UnsurMolekul string `form:"unsurMolekul" validate:"required"`
		Image        string `form:"image" validate:"required"`
		ImageUrl     string `form:"imageUrl" validate:"required"`
		Description  string `form:"description" validate:"required"`
		ListSenyawa  []Senyawa
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

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	var listSenyawa []utils.SenyawaRequest
	for _, senyawa := range payloadValidator.ListSenyawa {
		senyawa := utils.SenyawaRequest{
			Judul:     senyawa.Judul,
			Unsur:     senyawa.Unsur,
			Deskripsi: senyawa.Deskripsi,
		}

		if err := controller.validate.Struct(senyawa); err != nil {
			return c.JSON(400, err.Error())
		}

		listSenyawa = append(listSenyawa, senyawa)
	}

	magicCardReq := utils.MagicCardRequest{
		NamaMolekul:  payloadValidator.NamaMolekul,
		UnsurMolekul: payloadValidator.UnsurMolekul,
		Image:        payloadValidator.Image,
		ImageUrl:     payloadValidator.ImageUrl,
		Description:  payloadValidator.Description,
		ListSenyawa:  listSenyawa,
	}

	response := controller.magicCardService.UpdateMagicCard(uuid, magicCardReq, token, photoRequest)

	return c.JSON(response.StatusCode, response)
}

func (controller *MagicCardController) DeleteMagicCard(c echo.Context) error {
	id := c.Param("id")
	uuid, err := utils.ParseDataId(id)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.magicCardService.DeleteMagicCard(uuid, token)
	return c.JSON(response.StatusCode, response)
}

func (controller *MagicCardController) CreateSenyawaAndIntegrateToMagicCard(c echo.Context) error {

	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}
	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	id := c.Param("id")
	uuid, err := utils.ParseDataId(id)

	if err != nil {
		return c.JSON(400, err.Error())
	}

	type Senyawa struct {
		Judul     string `form:"judul" validate:"required"`
		Unsur     string `form:"unsur" validate:"required"`
		Deskripsi string `form:"deskripsi" validate:"required"`
	}

	payloadValidator := new(Senyawa)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	senyawaReq := utils.SenyawaRequest{
		Judul:     payloadValidator.Judul,
		Unsur:     payloadValidator.Unsur,
		Deskripsi: payloadValidator.Deskripsi,
	}

	response := controller.magicCardService.CreateSenyawaAndIntegrateToMagicCard(uuid, token, senyawaReq)

	return c.JSON(response.StatusCode, response)
}

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

type AuthController struct {
	authService services.AuthService
	validate    vl.Validate
}

func NewAuthController(db *gorm.DB) AuthController {
	service := services.NewAuthService(db)
	controller := AuthController{
		authService: service,
		validate:    *vl.New(),
	}
	return controller
}

func (controller *AuthController) Refresh(c echo.Context) error {
	type payload struct {
		RefreshToken string `json:"refreshToken"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.authService.Refresh(payloadValidator.RefreshToken)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) Login(c echo.Context) error {
	type payload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	response := controller.authService.Login(payloadValidator.Email, payloadValidator.Password)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) Register(c echo.Context) error {
	type payload struct {
		NamaLengkap          string    `json:"namaLengkap" validate:"required"`
		Email                string    `json:"email" validate:"required"`
		Password             string    `json:"password" validate:"required"`
		PasswordConfirmation string    `json:"passwordConfirmation" validate:"required,eqfield=Password"`
		Role                 string    `json:"role"`
		IdKelas              uuid.UUID `json:"idKelas"`
		IdSekolah            uuid.UUID `json:"idSekolah"`
	}

	payloadValidator := new(payload)

	if err := c.Bind(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	if err := controller.validate.Struct(payloadValidator); err != nil {
		return c.JSON(400, err.Error())
	}

	var regisUserPayload utils.UserRequest = utils.UserRequest{
		Fullname:             payloadValidator.NamaLengkap,
		IdKelas:              payloadValidator.IdKelas,
		IdSekolah:            payloadValidator.IdSekolah,
		Email:                payloadValidator.Email,
		Password:             payloadValidator.Password,
		PasswordConfirmation: payloadValidator.PasswordConfirmation,
		Role:                 payloadValidator.Role,
	}

	if regisUserPayload.Role == "guru" {
		regisUserPayload.IdKelas = uuid.Nil
		regisUserPayload.IdSekolah = uuid.Nil
	}

	if regisUserPayload.Role == "siswa" {
		idSekolah, _ := uuid.Parse("144358b8-1ce4-11ef-9b63-dead0d6128da")
		regisUserPayload.IdSekolah = idSekolah

		idKelas, _ := uuid.Parse("71f0e8e1-1ce4-11ef-9b63-dead0d6128da")
		regisUserPayload.IdKelas = idKelas
	}

	response := controller.authService.Register(regisUserPayload)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) GetUser(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")
	response := controller.authService.GetLoggedInUser(token)
	return c.JSON(response.StatusCode, response)
}

func (controller *AuthController) GetTeacher(c echo.Context) error {
	authorizationHeader := c.Request().Header.Get("Authorization")
	if authorizationHeader == "" || !strings.HasPrefix(authorizationHeader, "Bearer ") {
		return c.JSON(401, "Unauthorized")
	}

	token := strings.TrimPrefix(authorizationHeader, "Bearer ")

	response := controller.authService.GetTeacher(token)
	return c.JSON(response.StatusCode, response)
}

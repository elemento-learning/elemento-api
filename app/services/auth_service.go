package services

import (
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"
	"fmt"

	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	authRepo repositories.UserRepository
}

func (service *authService) Login(email, password string) utils.Response {
	var response utils.Response
	if email == "" || password == "" {
		response.StatusCode = 400
		response.Messages = "email dan password tidak boleh kosong"
		response.Data = nil
		return response
	}

	if !utils.IsEmailValid(email) {
		response.StatusCode = 400
		response.Messages = "Email kamu tidak valid"
		response.Data = nil
		return response
	}

	user, err := service.authRepo.GetUserByEmail(email)
	if err != nil {
		response.StatusCode = 401
		response.Messages = "Email kamu belum terdaftar"
		response.Data = nil
		return response
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		fmt.Print(!utils.CheckPasswordHash(password, user.Password))
		response.StatusCode = 401
		response.Messages = "Password kamu salah"
		response.Data = nil
		return response
	}
	AccessToken, err := utils.GenerateJWTAccessToken(user.IdUser, user.Fullname, user.Email, "elemento")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	refreshToken, err := utils.GenerateJWTRefreshToken(user.IdUser, user.Fullname, user.Email, "elemento")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  AccessToken,
		"refreshToken": refreshToken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

func (service *authService) Register(registerRequest utils.UserRequest) utils.Response {
	var response utils.Response
	if registerRequest.Fullname == "" || registerRequest.Email == "" || registerRequest.Password == "" || registerRequest.PasswordConfirmation == "" {
		response.StatusCode = 400
		response.Messages = "Semua field harus diisi"
		response.Data = nil
		return response
	}
	if !utils.IsEmailValid(registerRequest.Email) {
		response.StatusCode = 400
		response.Messages = "Email kamu tidak valid"
		response.Data = nil
		return response
	}

	user, err := service.authRepo.GetUserByEmail(registerRequest.Email)
	if err == nil {
		response.StatusCode = 400
		response.Messages = "Email sudah terdaftar"
		response.Data = nil
		return response
	}

	if registerRequest.Password != registerRequest.PasswordConfirmation {
		response.StatusCode = 400
		response.Messages = "Password dan konfirmasi password tidak sama"
		response.Data = nil
		return response
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Password hashing failed"
		response.Data = nil
		return response
	}
	userId := uuid.New()
	accessToken, err := utils.GenerateJWTAccessToken(userId, registerRequest.Fullname, registerRequest.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	refreshtoken, err := utils.GenerateJWTRefreshToken(userId, registerRequest.Fullname, registerRequest.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	user = models.User{
		IdUser:   userId,
		Fullname: registerRequest.Fullname,
		Email:    registerRequest.Email,
		Password: string(hashedPassword),
		Role:     registerRequest.Role,
	}

	err = service.authRepo.CreateNewUser(user)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "User creation failed"
		response.Data = user.IdUser
		return response
	}
	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshtoken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

func (service *authService) GetLoggedInUser(bearerToken string) utils.Response {
	var response utils.Response
	var firstname, lastname string
	id, err := utils.ParseDataId(bearerToken)
	if id != uuid.Nil && err == nil {
		user, err := service.authRepo.GetUserById(id)
		if err != nil {
			response.StatusCode = 500
			response.Messages = "User tidak ditemukan"
			response.Data = nil
			return response
		}

		response.StatusCode = 200
		response.Messages = "success"
		names := strings.Split(user.Fullname, " ")
		if len(names) == 1 {
			firstname = names[0]
			lastname = names[0]
		} else {
			firstname = names[0]
			lastname = names[len(names)-1]
		}
		if user.Role != "admin" {
			response.Data = map[string]interface{}{
				"idUser":    user.IdUser,
				"firstName": firstname,
				"lastName":  lastname,
				"email":     user.Email,
				"role":      user.Role,
			}
		} else {
			response.Data = map[string]interface{}{
				"idUser":    user.IdUser,
				"firstName": firstname,
				"lastName":  lastname,
				"email":     user.Email,
				"role":      user.Role,
			}
		}
		return response
	} else {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
}

func (service *authService) Refresh(refreshToken string) utils.Response {
	var response utils.Response
	userId, err := utils.ParseDataId(refreshToken)
	if userId == uuid.Nil || err != nil {
		response.StatusCode = 401
		response.Messages = "Invalid token"
		response.Data = nil
		return response
	}
	user, err := service.authRepo.GetUserById(userId)
	if err != nil {
		response.StatusCode = 401
		response.Messages = "User tidak ditemukan"
		response.Data = nil
		return response
	}
	AccessToken, err := utils.GenerateJWTAccessToken(user.IdUser, user.Fullname, user.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}
	refreshToken, err = utils.GenerateJWTRefreshToken(user.IdUser, user.Fullname, user.Email, "kalorize")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Token generation failed"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "success"
	response.Data = map[string]interface{}{
		"accessToken":  AccessToken,
		"refreshToken": refreshToken,
		"role":         user.Role,
		"userId":       user.IdUser,
	}
	return response
}

type AuthService interface {
	Login(username, password string) utils.Response
	Register(requestRegister utils.UserRequest) utils.Response
	GetLoggedInUser(bearerToken string) utils.Response
	Refresh(refreshToken string) utils.Response
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{
		authRepo: repositories.NewDBUserRepository(db),
	}
}

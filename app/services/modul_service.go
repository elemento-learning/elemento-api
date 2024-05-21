package services

import (
	"context"
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type modulService struct {
	modulRepo repositories.ModulRepository
	babRepo   repositories.BabRepository
}

func (service *modulService) CreateNewModul(modul utils.ModulRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	opt := option.WithCredentialsFile("config/config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal membuat aplikasi firebase",
			Data:       nil,
		}
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal membuat client firebase" + err.Error(),
			Data:       nil,
		}
	}
	formattedTitle := utils.FormatTitleFromFirebase(modul.Title)
	storagePath := "modul/" + formattedTitle + "/" + photoRequest.Handler.Filename
	reader := photoRequest.File

	bucket, err := client.Bucket("elemento-84e6b.appspot.com")
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal membuat bucket firebase",
			Data:       nil,
		}
	}

	wc := bucket.Object(storagePath).NewWriter(context.Background())
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	// Upload the file to Firebase Storage
	if _, err := io.Copy(wc, reader); err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to upload file to Firebase Storage",
			Data:       nil,
		}
	}

	newModul := models.Modul{
		Title:       modul.Title,
		Subtitle:    modul.Subtitle,
		YoutubeLink: modul.YoutubeLink,
	}

	newModul.ModulID = uuid.New()

	for _, bab := range newModul.Babs {
		err := service.babRepo.CreateNewBab(bab)
		if err != nil {
			return utils.Response{
				StatusCode: 500,
				Messages:   "Gagal membuat bab",
				Data:       nil,
			}
		}
	}

	// Close the writer after copying
	if err := wc.Close(); err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to close Firebase Storage writer" + err.Error(),
			Data:       nil,
		}
	}

	newModul.Image = photoRequest.Alias + "/" + storagePath
	newModul.ImageUrl = fmt.Sprintf("https://storage.googleapis.com/elemento-84e6b.appspot.com/%s", storagePath)

	var response utils.Response
	err = service.modulRepo.CreateNewModul(newModul)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat modul"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil membuat modul"
	response.Data = newModul
	return response
}

func (service *modulService) GetModul(bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}

	}

	var response utils.Response
	moduls, err := service.modulRepo.GetAllModul()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan modul"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil mendapatkan modul"
	response.Data = moduls
	return response
}

func (service *modulService) CreateBabAndIntegrateToModul(modulId uuid.UUID, bearerToken string, bab utils.BabRequest) utils.Response {

	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response

	newBab := models.Bab{
		Title:       bab.Title,
		Description: bab.Description,
		Task:        bab.Task,
	}

	newBab.TitleID = uuid.New()
	err := service.babRepo.CreateNewBab(newBab)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat bab"
		response.Data = nil
		return response
	}

	modul, err := service.modulRepo.GetModulById(modulId)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan modul"
		response.Data = nil
		return response
	}

	err = service.modulRepo.IntegrateBabToModul(modul, newBab)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mengintegrasikan bab ke modul" + err.Error()
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil membuat bab dan mengintegrasikannya ke modul"
	response.Data = bab
	return response
}

func (service *modulService) GetModulById(id uuid.UUID) utils.Response {
	var response utils.Response
	modul, err := service.modulRepo.RetrieveUpdatedModulWithAssociatedBab(id)

	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan modul"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil mendapatkan modul"
	response.Data = modul
	return response
}

type ModulService interface {
	CreateNewModul(modul utils.ModulRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response
	GetModulById(id uuid.UUID) utils.Response
	CreateBabAndIntegrateToModul(modulId uuid.UUID, bearerToken string, bab utils.BabRequest) utils.Response
	GetModul(bearerToken string) utils.Response
}

func NewModulService(db *gorm.DB) ModulService {
	return &modulService{modulRepo: repositories.NewModulRepository(db),
		babRepo: repositories.NewBabRepository(db)}
}

// Path: app/services/modul_service.go

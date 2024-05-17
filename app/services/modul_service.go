package services

import (
	"context"
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"
	"io"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"github.com/google/uuid"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type modulService struct {
	modulRepo repositories.ModulRepository
}

func (service *modulService) CreateNewModul(modul utils.ModulRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	newModul := models.Modul{
		Title:      modul.Title,
		Subtitle:   modul.Subtitle,
		IsComplete: modul.IsComplete,
	}

	for _, bab := range modul.Babs {
		newBab := models.Bab{
			Title:         bab.Title,
			Description:   bab.Description,
			Task:          bab.Task,
			ResultStudent: bab.ResultStudent,
		}
		newModul.Babs = append(newModul.Babs, newBab)
	}

	opt := option.WithCredentialsFile("credentials.json")
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
			Messages:   "Gagal membuat client firebase",
			Data:       nil,
		}
	}

	storagePath := "modul/" + modul.Title + "/" + photoRequest.Handler.Filename
	reader := photoRequest.File

	bucket, err := client.Bucket("gs://elemento-84e6b.appspot.com")
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

	// Close the writer after copying
	if err := wc.Close(); err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Failed to close Firebase Storage writer",
			Data:       nil,
		}
	}

	modul.Image = photoRequest.Alias + "/" + storagePath
	modul.ImageUrl = "https://firebasestorage.googleapis.com/v0/b/elemento-84e6b.appspot.com/o/" + storagePath + "?alt=media"

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
	response.Data = modul
	return response
}

func (service *modulService) GetModulById(id uuid.UUID) utils.Response {
	var response utils.Response
	modul, err := service.modulRepo.GetModulById(id)
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

func (service *modulService) GetModul() utils.Response {
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

type ModulService interface {
	CreateNewModul(modul utils.ModulRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response
	GetModulById(id uuid.UUID) utils.Response
	GetModul() utils.Response
}

func NewModulService(db *gorm.DB) ModulService {
	return &modulService{modulRepo: repositories.NewModulRepository(db)}
}

// Path: app/services/modul_service.go

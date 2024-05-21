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

type magicCardService struct {
	magicRepo   repositories.MagicCardRepository
	senyawaRepo repositories.SenyawaRepository
}

func (service *magicCardService) CreateMagicCard(magicCard utils.MagicCardRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response {

	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	if magicCard.NamaMolekul == "" || magicCard.UnsurMolekul == "" || magicCard.Description == "" {
		response.StatusCode = 400
		response.Messages = "Nama molekul, unsur molekul, dan deskripsi tidak boleh kosong"
		response.Data = nil
		return response
	}

	newMagicCard := models.MagicCard{
		NamaMolekul:  magicCard.NamaMolekul,
		UnsurMolekul: magicCard.UnsurMolekul,
		Image:        magicCard.Image,
		ImageUrl:     magicCard.ImageUrl,
		Description:  magicCard.Description,
	}

	for _, senyawa := range magicCard.ListSenyawa {
		newSenyawa := models.Senyawa{
			Judul:     senyawa.Judul,
			Unsur:     senyawa.Unsur,
			Deskripsi: senyawa.Deskripsi,
		}

		newMagicCard.ListSenyawa = append(newMagicCard.ListSenyawa, newSenyawa)
	}

	opt := option.WithCredentialsFile("config/config.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat aplikasi firebase"
		response.Data = nil
		return response
	}

	client, err := app.Storage(context.Background())
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat client firebase" + err.Error()
		response.Data = nil
		return response
	}

	bucket, err := client.Bucket("elemento-84e6b.appspot.com")
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat bucket firebase"
		response.Data = nil
		return response
	}

	formattedTitle := utils.FormatTitleFromFirebase(newMagicCard.NamaMolekul)
	storagePath := "magic_card/" + formattedTitle + "/" + photoRequest.Handler.Filename
	reader := photoRequest.File

	wc := bucket.Object(storagePath).NewWriter(context.Background())
	wc.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	if _, err := io.Copy(wc, reader); err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mengupload foto ke firebase storage"
		response.Data = nil
		return response
	}

	newMagicCard.Image = photoRequest.Alias + "/" + storagePath
	newMagicCard.ImageUrl = fmt.Sprintf("https://storage.googleapis.com/elemento-84e6b.appspot.com/%s", storagePath)

	err = service.magicRepo.CreateNewMagicCard(newMagicCard)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat kartu magic"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil membuat kartu magic oleh" + teacherId.String()
	response.Data = newMagicCard

	return response
}

func (service *magicCardService) GetMagicCardById(id uuid.UUID, bearerToken string) utils.Response {
	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	magicCard, err := service.magicRepo.GetMagicCardById(id)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan kartu magic"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil mendapatkan kartu magic oleh" + teacherId.String()
	response.Data = magicCard
	return response
}

func (service *magicCardService) GetAllMagicCard(bearerToken string) utils.Response {

	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	magicCards, err := service.magicRepo.GetAllMagicCard()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan kartu magic"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil mendapatkan kartu magic oleh" + teacherId.String()
	response.Data = magicCards
	return response

}

func (service *magicCardService) UpdateMagicCard(uuid uuid.UUID, magicCard utils.MagicCardRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response {
	var response utils.Response
	if magicCard.NamaMolekul == "" || magicCard.UnsurMolekul == "" || magicCard.Description == "" {
		response.StatusCode = 400
		response.Messages = "Nama molekul, unsur molekul, dan deskripsi tidak boleh kosong"
		response.Data = nil
	}

	return response
}

func (service *magicCardService) DeleteMagicCard(id uuid.UUID, bearerToken string) utils.Response {
	var response utils.Response

	err := service.magicRepo.DeleteMagicCard(id)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal menghapus kartu magic"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil menghapus kartu magic"
	response.Data = nil
	return response

}

func (service *magicCardService) CreateSenyawaAndIntegrateToMagicCard(magicCardId uuid.UUID, bearerToken string, senyawa utils.SenyawaRequest) utils.Response {
	var response utils.Response
	if senyawa.Judul == "" || senyawa.Unsur == "" || senyawa.Deskripsi == "" {
		response.StatusCode = 400
		response.Messages = "Judul, unsur, dan deskripsi tidak boleh kosong"
		response.Data = nil
		return response
	}

	newSenyawa := models.Senyawa{
		Judul:     senyawa.Judul,
		Unsur:     senyawa.Unsur,
		Deskripsi: senyawa.Deskripsi,
	}

	err := service.senyawaRepo.CreateNewSenyawa(newSenyawa)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal membuat senyawa"
		response.Data = nil
		return response
	}

	magicCard, err := service.magicRepo.GetMagicCardById(magicCardId)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan kartu magic"
		response.Data = nil
		return response
	}

	err = service.magicRepo.IntegrateSenyawaToMagicCard(magicCard, newSenyawa)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mengintegrasikan senyawa ke kartu magic"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil membuat senyawa dan mengintegrasikannya ke kartu magic"
	response.Data = senyawa
	return response
}

type MagicCardService interface {
	CreateMagicCard(magicCard utils.MagicCardRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response
	GetMagicCardById(id uuid.UUID, bearerToken string) utils.Response
	GetAllMagicCard(bearerToken string) utils.Response
	CreateSenyawaAndIntegrateToMagicCard(magicCardId uuid.UUID, bearerToken string, senyawa utils.SenyawaRequest) utils.Response
	UpdateMagicCard(uuid uuid.UUID, magicCard utils.MagicCardRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response
	DeleteMagicCard(id uuid.UUID, bearerToken string) utils.Response
}

func NewMagicCardService(db *gorm.DB) MagicCardService {
	return &magicCardService{magicRepo: repositories.NewMagicCardRepository(db)}
}

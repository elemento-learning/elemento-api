package services

import (
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type magicCardService struct {
	magicRepo repositories.MagicCardRepository
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

func (service *magicCardService) UpdateMagicCard(magicCard utils.MagicCardRequest, bearerToken string) utils.Response {
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

type MagicCardService interface {
	CreateMagicCard(magicCard utils.MagicCardRequest, bearerToken string, photoRequest utils.UploadedPhoto) utils.Response
	GetMagicCardById(id uuid.UUID, bearerToken string) utils.Response
	GetAllMagicCard(bearerToken string) utils.Response
	UpdateMagicCard(magicCard utils.MagicCardRequest, bearerToken string) utils.Response
	DeleteMagicCard(id uuid.UUID, bearerToken string) utils.Response
}

func NewMagicCardService(db *gorm.DB) MagicCardService {
	return &magicCardService{magicRepo: repositories.NewMagicCardRepository(db)}
}

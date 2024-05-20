package services

import (
	"elemento-api/app/models"
	"elemento-api/app/repositories"
	"elemento-api/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SchoolService is a struct to define the service for school
type schoolService struct {
	schoolRepo repositories.SchoolRepository
	classRepo  repositories.ClassRepository
}

// CreateNewSchool is a function to create new school
func (s *schoolService) CreateNewSchool(school utils.SchoolRequest, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	newSchool := models.School{
		Name:     school.Name,
		Location: school.Location,
	}

	err = s.schoolRepo.CreateNewSchool(newSchool)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal membuat sekolah" + err.Error(),
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil membuat sekolah" + teacherId.String(),
		Data:       newSchool,
	}
}

// GetSchoolById is a function to get school by id
func (s *schoolService) GetSchoolById(uuid uuid.UUID, bearerToken string) utils.Response {
	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	school, err := s.schoolRepo.GetSchoolById(uuid)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan sekolah"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil mendapatkan sekolah oleh" + teacherId.String()
	response.Data = school
	return response
}

// GetAllSchool is a function to get all school
func (s *schoolService) GetAllSchool(bearerToken string) utils.Response {
	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	schools, err := s.schoolRepo.GetAllSchool()
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan sekolah"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil mendapatkan sekolah oleh" + teacherId.String()
	response.Data = schools
	return response
}

// UpdateSchool is a function to update school
func (s *schoolService) UpdateSchool(school utils.SchoolRequest, bearerToken string) utils.Response {
	if bearerToken == "" {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	newSchool := models.School{
		Name:     school.Name,
		Location: school.Location,
	}

	err = s.schoolRepo.UpdateSchool(newSchool)
	if err != nil {
		return utils.Response{
			StatusCode: 500,
			Messages:   "Gagal mengupdate sekolah" + err.Error(),
			Data:       nil,
		}
	}

	return utils.Response{
		StatusCode: 200,
		Messages:   "Berhasil mengupdate sekolah" + teacherId.String(),
		Data:       newSchool,
	}
}

// DeleteSchool is a function to delete school
func (s *schoolService) DeleteSchool(uuid uuid.UUID, bearerToken string) utils.Response {
	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	err = s.schoolRepo.DeleteSchool(uuid)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal menghapus sekolah"
		response.Data = nil
		return response
	}
	response.StatusCode = 200
	response.Messages = "Berhasil menghapus sekolah oleh" + teacherId.String()
	response.Data = nil
	return response
}

// IntegrateClassToSchool is a function to integrate class to school
func (s *schoolService) IntegrateClassToSchool(schoolId uuid.UUID, class utils.ClassRequest, bearerToken string) utils.Response {
	teacherId, err := utils.ParseDataId(bearerToken)

	if err != nil {
		return utils.Response{
			StatusCode: 401,
			Messages:   "Unauthorized",
			Data:       nil,
		}
	}

	var response utils.Response
	school, err := s.schoolRepo.GetSchoolById(schoolId)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mendapatkan sekolah"
		response.Data = nil
		return response
	}

	newClass := models.Class{
		Name:     class.Name,
		Location: class.Location,
	}

	err = s.schoolRepo.IntegrateClassToSchool(school, newClass)
	if err != nil {
		response.StatusCode = 500
		response.Messages = "Gagal mengintegrasikan kelas ke sekolah"
		response.Data = nil
		return response
	}

	response.StatusCode = 200
	response.Messages = "Berhasil mengintegrasikan kelas ke sekolah oleh" + teacherId.String()
	response.Data = nil
	return response
}

type SchoolService interface {
	CreateNewSchool(school utils.SchoolRequest, bearerToken string) utils.Response
	GetSchoolById(uuid uuid.UUID, bearerToken string) utils.Response
	GetAllSchool(bearerToken string) utils.Response
	UpdateSchool(school utils.SchoolRequest, bearerToken string) utils.Response
	DeleteSchool(uuid uuid.UUID, bearerToken string) utils.Response
	IntegrateClassToSchool(schoolId uuid.UUID, class utils.ClassRequest, bearerToken string) utils.Response
}

// NewSchoolService is a function to create new school service
func NewSchoolService(db *gorm.DB) SchoolService {
	return &schoolService{
		schoolRepo: repositories.NewSchoolRepository(db),
		classRepo:  repositories.NewClassRepository(db),
	}
}

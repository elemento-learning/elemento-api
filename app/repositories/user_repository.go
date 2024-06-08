package repositories

import (
	"elemento-api/app/models"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type dbUser struct {
	Conn *gorm.DB
}

func (dbAuth *dbUser) GetToken() string {
	return "token"
}

func (db *dbUser) GetUserById(id uuid.UUID) (user models.User, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting User by ID: %v", err)
		}
	}()
	err = db.Conn.Where("id_user = ?", id).First(&user).Error
	return user, err
}

func (db *dbUser) CreateNewUser(user models.User) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error creating new User: %v", err)
		}
	}()
	err = db.Conn.Create(&user).Error
	return err
}

func (db *dbUser) GetUserByUsername(username string) (user models.User, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting User by username: %v", err)
		}
	}()
	err = db.Conn.Where("full_name = ?", username).First(&user).Error
	return user, err
}

func (db *dbUser) GetUserByEmail(email string) (user models.User, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting User by email: %v", err)
		}
	}()
	err = db.Conn.Where("email = ?", email).First(&user).Error
	return user, err
}

func (db *dbUser) FindReferalCodeIfExist(code string) (exists bool) {
	var user models.User
	err := db.Conn.Where("referal_code = ?", code).First(&user).Error
	if err != nil {
		log.Printf("Error finding referral code: %v", err)
		return false
	}
	return true
}

func (db *dbUser) UpdateUser(user models.User) (err error) {
	defer func() {
		if err != nil {
			log.Printf("Error updating User: %v", err)
		}
	}()
	err = db.Conn.Save(&user).Error
	return err
}

func (db *dbUser) GetTeacher() (teachers []models.User, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting Teachers: %v", err)
		}
	}()
	err = db.Conn.Where("role = ?", "guru").Find(&teachers).Error
	return teachers, err
}

func (db *dbUser) GetTeacherById(id uuid.UUID) (teacher models.User, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting Teacher by ID: %v", err)
		}
	}()
	err = db.Conn.Where("id_user = ? AND role = ?", id, "guru").First(&teacher).Error
	return teacher, err
}

func (db *dbUser) GetStudentById(id uuid.UUID) (student models.User, err error) {
	defer func() {
		if err != nil {
			log.Printf("Error getting Student by ID: %v", err)
		}
	}()
	err = db.Conn.Where("id_user = ? AND role = ?", id, "siswa").First(&student).Error
	return student, err
}

type UserRepository interface {
	GetToken() string
	GetStudentById(id uuid.UUID) (models.User, error)
	GetTeacherById(id uuid.UUID) (models.User, error)
	CreateNewUser(user models.User) error
	GetUserByUsername(username string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	FindReferalCodeIfExist(code string) bool
	UpdateUser(user models.User) error
	GetTeacher() ([]models.User, error)
	GetUserById(id uuid.UUID) (models.User, error)
}

func NewDBUserRepository(conn *gorm.DB) *dbUser {
	return &dbUser{Conn: conn}
}

package utils

import "github.com/google/uuid"

type UserRequest struct {
	IdUser               uuid.UUID `json:"id_user"`
	Fullname             string    `json:"fullname"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	PasswordConfirmation string    `json:"password_confirmation"`
	Role                 string    `json:"role"`
	IdKelas              uuid.UUID `json:"id_kelas"`
	IdSekolah            uuid.UUID `json:"id_sekolah"`
}

func ValidateAndAssign(target *string, source string) {
	if source != "" {
		*target = source
	}
}

func ValidateAndAssignInt(target *int, source *int) {
	if source != nil {
		*target = *source
	}
}

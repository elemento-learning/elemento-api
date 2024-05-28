package models

import "github.com/google/uuid"

type TargetKalori struct {
}

type User struct {
	IdUser    uuid.UUID `json:"id_user" gorm:"column:id_user;primary_key;type:char(36);"`
	Fullname  string    `json:"fullname" gorm:"column:full_name;type:varchar(255);"`
	Email     string    `json:"email" gorm:"column:email;type:varchar(255);"`
	Password  string    `json:"password" gorm:"column:password;type:varchar(255);"`
	Role      string    `json:"role" gorm:"column:role;type:varchar(20);"`
	IdKelas   uuid.UUID `json:"id_kelas" gorm:"column:id_kelas;type:char(36);"`
	IdSekolah uuid.UUID `json:"id_sekolah" gorm:"column:id_sekolah;type:char(36);"`
}

func (u *User) TableName() string {
	return "users"
}

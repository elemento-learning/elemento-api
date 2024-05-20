package models

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type School struct {
	gorm.Model
	SchoolID uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	Name     string    `gorm:"column:name;type:varchar(255);"`
	Location string    `gorm:"column:location;type:varchar(255);"`
	Class    []Class   `gorm:"foreignKey:SchoolID"`
}

type Class struct {
	gorm.Model
	ClassID  uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	Name     string    `gorm:"column:name;type:varchar(255);"`
	Location string    `gorm:"column:location;type:varchar(255);"`
	SchoolID uuid.UUID `gorm:"column:school_id;type:char(36);"`
}

func (u *School) TableName() string {
	return "schools"
}

func (u *Class) TableName() string {
	return "classes"
}

// Path: app/models/school.go

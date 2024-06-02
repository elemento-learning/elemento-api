package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeedBack struct {
	gorm.Model
	TeacherID  uuid.UUID `gorm:"column:teacher_id;type:char(36);"`
	FeedBackID uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	StudentID  uuid.UUID `gorm:"column:student_id;type:char(36);"`
	FeedBack   string    `gorm:"column:feedback;type:text;"`
	Category   string    `gorm:"column:category;type:varchar(255);"`
}

func (FeedBack) TableName() string {
	return "feedbacks"
}

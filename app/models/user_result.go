package models

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type UserResult struct {
	gorm.Model
	UserResultID uuid.UUID    `gorm:"column:id;type:char(36);primary_key;"`
	UserID       uuid.UUID    `gorm:"column:user_id;type:char(36);"`
	QuizID       uuid.UUID    `gorm:"column:quiz_id;type:char(36);"`
	CountAnswer  int          `gorm:"column:count_answer;type:int;"`
	Score        int          `gorm:"column:score;type:int;"`
	Answer       []UserAnswer `gorm:"foreignKey:UserAnswerID"`
}

type UserAnswer struct {
	gorm.Model
	UserAnswerID       uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	TitleQuestion      string    `gorm:"column:title_question;type:varchar(255);"`
	UserAnswerTitle    string    `gorm:"column:user_answer_title;type:varchar(255);"`
	UserAnswerSubtitle string    `gorm:"column:user_answer_subtitle;type:varchar(255);"`
}

func (UserResult) TableName() string {
	return "user_result"
}

func (UserAnswer) TableName() string {
	return "user_answer"
}

package models

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	QuizID   uuid.UUID  `gorm:"column:id;type:char(36);primary_key;"`
	Title    string     `gorm:"column:title;type:varchar(255);"`
	Status   string     `gorm:"column:status;type:varchar(255);"`
	Question []Question `gorm:"foreignKey:QuizID;references:QuizID"`
}

type Question struct {
	gorm.Model
	QuestionID uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	QuizID     uuid.UUID `gorm:"column:quiz_id;type:char(36);"`
	Question   string    `gorm:"column:question;type:varchar(255);"`
	Answer     []Answer  `gorm:"foreignKey:QuestionID;references:QuestionID"`
	AnswerQuiz uuid.UUID `gorm:"column:answer_quiz;type:char(36);"`
}

type Answer struct {
	gorm.Model
	AnswerID       uuid.UUID `gorm:"column:id;type:char(36);primary_key;"`
	AnswerTitle    string    `gorm:"column:answer_title;type:varchar(255);"`
	AnswerSubtitle string    `gorm:"column:answer_subtitle;type:varchar(255);"`
}

func (Quiz) TableName() string {
	return "quiz"
}

func (Question) TableName() string {
	return "question"
}

func (Answer) TableName() string {
	return "answer"
}

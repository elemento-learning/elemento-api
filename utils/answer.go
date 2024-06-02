package utils

import "github.com/google/uuid"

type AnswerRequest struct {
	AnswerID       uuid.UUID
	AnswerTitle    string
	AnswerSubtitle string
	QuestionID     uuid.UUID
}

type UserAnswerRequest struct {
	QuestionID uuid.UUID `json:"question_id"`
	AnswerID   uuid.UUID `json:"answer_id"`
}

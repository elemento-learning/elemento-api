package utils

import "github.com/google/uuid"

type QuestionRequest struct {
	Question       string    `json:"question"`
	QuestionAnswer uuid.UUID `json:"question_answer"`
}

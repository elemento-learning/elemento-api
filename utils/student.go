package utils

import "github.com/google/uuid"

type StudentScore struct {
	StudentID uuid.UUID `json:"student_id"`
	Score     int       `json:"score"`
	Name      string    `json:"name"`
}

package utils

import "github.com/google/uuid"

type StudentScore struct {
	ID        uuid.UUID `json:"id"`
	StudentID uuid.UUID `json:"student_id"`
	Score     int       `json:"score"`
	Name      string    `json:"name"`
}

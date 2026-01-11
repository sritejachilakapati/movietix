package domain

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID            uuid.UUID  `json:"id"`
	Title         string     `json:"title"`
	Description   string     `json:"description"`
	Language      string     `json:"language"`
	DurationMin   int        `json:"duration_min"`
	Certification *string    `json:"certification,omitempty"`
	Rating        *float32   `json:"rating,omitempty"`
	ReleaseDate   *time.Time `json:"release_date,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}

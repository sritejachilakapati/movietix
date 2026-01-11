package domain

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	ID             uuid.UUID  `json:"id"`
	Title          string     `json:"title"`
	Synopsis       string     `json:"synopsis"`
	LanguageCode   string     `json:"languageCode"`
	RuntimeMinutes int        `json:"runtimeMinutes"`
	PosterURL      *string    `json:"posterUrl,omitempty"`
	TrailerURL     *string    `json:"trailerUrl,omitempty"`
	Rating         *float32   `json:"rating,omitempty"`
	ReleaseDate    *time.Time `json:"releaseDate,omitempty"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Screen struct {
	ID         uuid.UUID `json:"id"`
	TheaterID  uuid.UUID `json:"theaterId"`
	Name       string    `json:"name"`
	TotalSeats int       `json:"totalSeats"`
	Formats    []Format  `json:"formats,omitempty"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
type Format string

const (
	Format2D         Format = "2d"
	Format3D         Format = "3d"
	FormatIMAX       Format = "imax"
	FormatDolbyAtmos Format = "dolby_atmos"
	FormatDolby71    Format = "dolby_7_1"
	Format4KAtmos    Format = "4k_dolby_atmos"
)

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Screen struct {
	ID         uuid.UUID      `json:"id"`
	TheaterID  uuid.UUID      `json:"theaterId"`
	Name       string         `json:"name"`
	TotalSeats int            `json:"totalSeats"`
	Formats    []ScreenFormat `json:"formats,omitempty"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
}
type ScreenFormat string

const (
	Format2D         ScreenFormat = "2d"
	Format3D         ScreenFormat = "3d"
	FormatIMAX       ScreenFormat = "imax"
	FormatDolbyAtmos ScreenFormat = "dolby_atmos"
	FormatDolby71    ScreenFormat = "dolby_7_1"
	Format4K         ScreenFormat = "4k"
)

package domain

import (
	"time"

	"github.com/google/uuid"
)

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Theater struct {
	ID           uuid.UUID     `json:"id"`
	Name         string        `json:"name"`
	CityCode     string        `json:"cityCode"`
	Address      string        `json:"address"`
	Location     Location      `json:"location,omitempty"`
	ContactPhone *string       `json:"contactPhone,omitempty"`
	ContactEmail *string       `json:"contactEmail,omitempty"`
	Status       TheaterStatus `json:"status"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
}

type TheaterStatus string

const (
	TheaterStatusActive   TheaterStatus = "active"
	TheaterStatusInactive TheaterStatus = "inactive"
)

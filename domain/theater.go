package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/twpayne/go-geom"
)

type Theater struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	CityCode  string        `json:"city_code"`
	Address   string        `json:"address"`
	Location  *geom.Point   `json:"location,omitempty"`
	Status    TheaterStatus `json:"status"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type TheaterStatus string

const (
	TheaterStatusActive   TheaterStatus = "active"
	TheaterStatusInactive TheaterStatus = "inactive"
)

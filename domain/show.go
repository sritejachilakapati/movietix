package domain

import (
	"time"

	"github.com/google/uuid"
)

type Show struct {
	ID        uuid.UUID      `json:"id"`
	MovieID   uuid.UUID      `json:"movieId"`
	ScreenID  uuid.UUID      `json:"screenId"`
	StartTime time.Time      `json:"startTime"`
	Status    ShowStatus     `json:"status"`
	Formats   []ScreenFormat `json:"formats,omitempty"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

type ShowSeatPricing struct {
	ShowID   uuid.UUID `json:"showId"`
	SeatType SeatType  `json:"seatType"`
	Price    int64     `json:"price"`
}

type ShowStatus string

const (
	ShowStatusScheduled ShowStatus = "scheduled"
	ShowStatusRunning   ShowStatus = "running"
	ShowStatusCancelled ShowStatus = "cancelled"
	ShowStatusCompleted ShowStatus = "completed"
)

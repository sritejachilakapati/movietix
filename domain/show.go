package domain

import (
	"time"

	"github.com/google/uuid"
)

type Show struct {
	ID        uuid.UUID      `json:"id"`
	MovieID   uuid.UUID      `json:"movie_id"`
	ScreenID  uuid.UUID      `json:"screen_id"`
	StartTime time.Time      `json:"start_time"`
	EndTime   time.Time      `json:"end_time"`
	Status    ShowStatus     `json:"status"`
	Formats   []ScreenFormat `json:"formats,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type ShowSeatPricing struct {
	ShowID   uuid.UUID `json:"show_id"`
	SeatType SeatType  `json:"seat_type"`
	Price    int64     `json:"price"`
}

type ShowStatus string

const (
	ShowStatusScheduled ShowStatus = "scheduled"
	ShowStatusRunning   ShowStatus = "running"
	ShowStatusCancelled ShowStatus = "cancelled"
	ShowStatusCompleted ShowStatus = "completed"
)

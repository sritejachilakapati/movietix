package domain

import "github.com/google/uuid"

type Seat struct {
	ID       uuid.UUID  `json:"id"`
	ScreenID uuid.UUID  `json:"screenId"`
	Row      string     `json:"row"`
	Number   int        `json:"number"`
	Type     SeatType   `json:"type"`
	Status   SeatStatus `json:"status"`
}

type SeatType string

const (
	SeatTypeRegular  SeatType = "regular"
	SeatTypePrime    SeatType = "prime"
	SeatTypeRecliner SeatType = "recliner"
)

type SeatStatus string

const (
	SeatStatusActive   SeatStatus = "active"
	SeatStatusInactive SeatStatus = "inactive"
)

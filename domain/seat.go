package domain

import "github.com/google/uuid"

type Seat struct {
	ID       uuid.UUID `json:"id"`
	ScreenID uuid.UUID `json:"screen_id"`
	Row      string    `json:"row"`
	Number   int       `json:"number"`
	Type     SeatType  `json:"type"`
}

type SeatType string

const (
	SeatTypeRegular  SeatType = "regular"
	SeatTypePremium  SeatType = "premium"
	SeatTypeRecliner SeatType = "recliner"
)

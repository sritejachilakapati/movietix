package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        uuid.UUID     `json:"id"`
	UserID    uuid.UUID     `json:"user_id"`
	ShowID    uuid.UUID     `json:"show_id"`
	Status    BookingStatus `json:"status"`
	TotalCost int64         `json:"total_cost"`
	Items     []BookingItem `json:"items"`
	CreatedAt time.Time     `json:"created_at"`
}

type BookingItem struct {
	ID        uuid.UUID `json:"id"`
	BookingID uuid.UUID `json:"booking_id"`
	SeatID    uuid.UUID `json:"seat_id"`
	Price     int64     `json:"price"`
}

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

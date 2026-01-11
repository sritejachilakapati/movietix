package domain

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID          uuid.UUID     `json:"id"`
	UserID      uuid.UUID     `json:"userId"`
	ShowID      uuid.UUID     `json:"showId"`
	Status      BookingStatus `json:"status"`
	TotalAmount int64         `json:"totalAmount"`
	Items       []BookingItem `json:"items"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
}

type BookingItem struct {
	ID        uuid.UUID `json:"id"`
	BookingID uuid.UUID `json:"bookingId"`
	SeatID    uuid.UUID `json:"seatId"`
	Price     int64     `json:"price"`
}

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

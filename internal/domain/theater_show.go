package domain

import (
	"time"

	"github.com/google/uuid"
)

type MovieTheaterShows struct {
	MovieID    uuid.UUID      `json:"movieId"`
	MovieTitle string         `json:"movieTitle"`
	Theaters   []TheaterShows `json:"theaters"`
}

type TheaterShows struct {
	TheaterID   uuid.UUID  `json:"theaterId"`
	TheaterName string     `json:"theaterName"`
	Address     string     `json:"address"`
	Location    Location   `json:"location"`
	Shows       []Showtime `json:"shows"`
}

type Showtime struct {
	ShowID    uuid.UUID `json:"showId"`
	StartTime time.Time `json:"startTime"`
	Format    Format    `json:"format"`
}

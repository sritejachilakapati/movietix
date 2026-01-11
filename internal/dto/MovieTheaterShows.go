package dto

import (
	"github.com/sritejachilakapati/movietix/internal/domain"
)

type MovieTheaterShows struct {
	domain.MovieTheaterShows
	NextOffset *int32 `json:"nextOffset"`
}

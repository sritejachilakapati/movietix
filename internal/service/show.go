package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/sritejachilakapati/movietix/internal/domain"
	"github.com/sritejachilakapati/movietix/internal/dto"
	"github.com/sritejachilakapati/movietix/internal/repository"
)

type ShowService interface {
	GetShowsByMovieAndCity(
		ctx context.Context,
		limit int32,
		offset int32,
		movieId uuid.UUID,
		cityCode string,
	) (dto.MovieTheaterShows, error)
}

type showService struct {
	queries *repository.Queries
}

func (s *showService) GetShowsByMovieAndCity(ctx context.Context, limit int32, offset int32, movieId uuid.UUID, cityCode string) (dto.MovieTheaterShows, error) {
	params := repository.GetShowsByMovieAndCityParams{
		Limit:    limit,
		Offset:   offset,
		MovieID:  movieId,
		CityCode: cityCode,
	}

	emptyResult := dto.MovieTheaterShows{}

	dbRows, err := s.queries.GetShowsByMovieAndCity(ctx, params)
	if err != nil {
		return emptyResult, err
	}

	if len(dbRows) == 0 {
		return emptyResult, nil
	}

	theaterMap := make(map[uuid.UUID]*domain.TheaterShows)
	for _, row := range dbRows {
		t, exists := theaterMap[row.TheaterID]
		if !exists {
			t = &domain.TheaterShows{
				TheaterID:   row.TheaterID,
				TheaterName: row.TheaterName,
				Address:     row.Address,
				Location:    locationFromGeomPoint(row.Location),
			}
			theaterMap[row.TheaterID] = t
		}

		t.Shows = append(t.Shows, domain.Showtime{
			ShowID:    row.ShowID,
			StartTime: row.StartTime,
			Format:    domain.Format(row.FormatName),
		})
	}

	theaters := make([]domain.TheaterShows, 0, len(theaterMap))
	for _, t := range theaterMap {
		theaters = append(theaters, *t)
	}

	nextOffset := calcNextOffset(limit, offset, len(dbRows))

	return dto.MovieTheaterShows{
		MovieTheaterShows: domain.MovieTheaterShows{
			MovieID:    movieId,
			MovieTitle: dbRows[0].MovieTitle,
			Theaters:   theaters,
		},
		NextOffset: nextOffset,
	}, nil

}

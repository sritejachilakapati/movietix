package service

import (
	"context"

	"github.com/sritejachilakapati/movietix/internal/domain"
	"github.com/sritejachilakapati/movietix/internal/dto"
	"github.com/sritejachilakapati/movietix/internal/repository"
)

type MovieService interface {
	GetMoviesByCity(
		context context.Context,
		cityCode string,
		limit int32,
		offset int32,
	) (dto.PageResult[domain.Movie], error)
}

type movieService struct {
	queries *repository.Queries
}

func (m *movieService) GetMoviesByCity(ctx context.Context, cityCode string, limit int32, offset int32) (dto.PageResult[domain.Movie], error) {
	params := repository.GetMoviesByCityParams{
		CityCode: cityCode,
		Limit:    limit,
		Offset:   offset,
	}
	dbMoviesRow, err := m.queries.GetMoviesByCity(ctx, params)
	if err != nil {
		return dto.PageResult[domain.Movie]{}, err
	}
	movies := moviesFromDB(dbMoviesRow)
	nextOffset := calcNextOffset(limit, offset, len(dbMoviesRow))

	return dto.PageResult[domain.Movie]{
		Items:      movies,
		NextOffset: nextOffset,
	}, nil
}

func NewMovieService(queries *repository.Queries) MovieService {
	return &movieService{
		queries: queries,
	}
}

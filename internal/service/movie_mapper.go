package service

import (
	"github.com/sritejachilakapati/movietix/internal/domain"
	"github.com/sritejachilakapati/movietix/internal/repository"
)

func movieFromDB(dbMovie repository.Movie) domain.Movie {
	var rating *float32

	if dbMovie.Rating.Valid {
		if f64Val, err := dbMovie.Rating.Float64Value(); err == nil {
			r := float32(f64Val.Float64)
			rating = &r
		}
	}

	certification := domain.CertificationUnrated
	if dbMovie.Certification != nil {
		switch *dbMovie.Certification {
		case "U":
			certification = domain.CertificationU
			break
		case "A":
			certification = domain.CertificationA
			break
		case "U/A":
			certification = domain.CertificationUA
			break
		default:
			certification = domain.CertificationUnknown
		}
	}

	return domain.Movie{
		ID:             dbMovie.ID,
		Title:          dbMovie.Title,
		Synopsis:       dbMovie.Synopsis,
		LanguageCode:   dbMovie.LanguageCode,
		RuntimeMinutes: dbMovie.RuntimeMinutes,
		Certification:  certification,
		PosterURL:      dbMovie.PosterUrl,
		TrailerURL:     dbMovie.TrailerUrl,
		Rating:         rating,
		ReleaseDate:    dbMovie.ReleaseDate,
		CreatedAt:      dbMovie.CreatedAt,
	}
}

func moviesFromDB(dbMovies []repository.Movie) []domain.Movie {
	movies := make([]domain.Movie, 0)
	for _, row := range dbMovies {
		movies = append(movies, movieFromDB(row))
	}

	return movies
}

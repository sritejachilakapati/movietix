package service

import (
	"github.com/sritejachilakapati/movietix/internal/domain"
	"github.com/twpayne/go-geom"
)

func calcNextOffset(limit int32, offset int32, resultsCount int) *int32 {
	if int32(resultsCount) < limit {
		return nil
	}
	newOffset := offset + limit
	return &newOffset
}

func locationFromGeomPoint(geom geom.Point) domain.Location {
	return domain.Location{
		Latitude:  geom.X(),
		Longitude: geom.Y(),
	}
}

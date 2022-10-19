package service

import (
	"context"

	sugar "github.com/kirandesimone/mflix/movie-app/internal"
	"github.com/kirandesimone/mflix/movie-app/internal/core/models"
	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

type Service struct {
	db ports.Db
}

func New(store ports.Db) *Service {
	return &Service{db: store}
}

func (s *Service) GetMovies(ctx context.Context) []*models.Movie {
	var movies []*models.Movie
	err := s.db.FindAll(ctx, &movies)
	if err != nil {
		sugar.Logger.Fatalf("SERVICE - %s", err)
	}
	sugar.Logger.Info("SERVICE - called getMovies")

	return movies

}

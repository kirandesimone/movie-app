package service

import (
	"context"
	"log"

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
		log.Fatal(err)
	}

	return movies

}

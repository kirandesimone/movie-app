package service

import (
	"context"

	sugar "github.com/kirandesimone/mflix/movie-app/internal"
	"github.com/kirandesimone/mflix/movie-app/internal/core/models"
	"github.com/kirandesimone/mflix/movie-app/internal/helpers"
	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

type Service struct {
	db ports.Db
}

func New(store ports.Db) *Service {
	return &Service{db: store}
}

func (s *Service) GetMovies() []*models.Movie {
	var movies []*models.Movie
	err := s.db.FindAll(context.Background(), &movies)
	if err != nil {
		sugar.Logger.Fatalf("SERVICE - %s", err)
	}
	sugar.Logger.Info("SERVICE - called getMovies")

	return movies

}

func (s *Service) GetTopRatedMovies() []*models.Movie {
	var topRated []*models.Movie

	err := s.db.TopRatedMovies(context.Background(), &topRated)
	if err != nil {
		sugar.Logger.Fatalf("SERVICE(TOP RATED MOVIES) - %s", err)
	}
	for i := 0; i < len(topRated); i++ {
		topRated[i].Poster = helpers.ValidateImg(topRated[i].Poster)
	}
	sugar.Logger.Info("SERVICE - Batching Top Rated Movies...Gathering 20")
	return topRated
}

func (s *Service) GetGenreMovies(genre string) []*models.Movie {
	var genreMovies []*models.Movie
	err := s.db.GenreMovies(context.Background(), genre, &genreMovies)
	if err != nil {
		sugar.Logger.Fatalf("SERIVEC(%s MOVIES) - %s", genre, err)
	}
	for i := 0; i < len(genreMovies); i++ {
		genreMovies[i].Poster = helpers.ValidateImg(genreMovies[i].Poster)
	}
	sugar.Logger.Infof("SERVICE - Batching %s Movies...Gathering 20", genre)
	return genreMovies
}

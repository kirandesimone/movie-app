package ports

import (
	"github.com/kirandesimone/mflix/movie-app/internal/core/models"
)

type Api interface {
	GetMovies() []*models.Movie
	GetTopRatedMovies() []*models.Movie
	GetGenreMovies(genre string) []*models.Movie
}

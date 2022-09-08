package ports

import (
	"context"

	"github.com/kirandesimone/mflix/movie-app/internal/core/models"
)

type Api interface {
	GetMovies(ctx context.Context) []*models.Movie
}

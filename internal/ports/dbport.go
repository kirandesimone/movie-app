package ports

import (
	"context"
)

type Db interface {
	Disconnect() error
	FindAll(ctx context.Context, results interface{}) error
	TopRatedMovies(ctx context.Context, results interface{}) error
	GenreMovies(ctx context.Context, genre string, results interface{}) error
}

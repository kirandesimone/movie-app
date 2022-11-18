package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

func HttpRouter(app *fiber.App, handler ports.Handler) {
	app.Get("/api/v1/movies", handler.GetMovies)
	app.Get("/api/v1/top/rated/movies", handler.GetTopRatedMovies)
	app.Get("/api/v1/genres/drama", handler.GetDramaMovies)
	app.Get("/api/v1/genres/action", handler.GetActionMovies)
}

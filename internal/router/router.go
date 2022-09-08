package router

import (
	"github.com/gorilla/mux"
	"github.com/kirandesimone/mflix/movie-app/internal/infrastructure/http"
	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

func HttpRouter(r *mux.Router, app ports.Api) {
	handler := http.NewHandler(app)

	r.HandleFunc("/api/v1/movies", handler.GetMovies)
}

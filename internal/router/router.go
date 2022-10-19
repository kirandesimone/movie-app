package router

import (
	"github.com/gorilla/mux"
	"github.com/kirandesimone/mflix/movie-app/internal/infrastructure/http"
	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

func HttpRouter(r *mux.Router, service ports.Api) {
	handler := http.NewHandler(service)

	r.HandleFunc("/api/v1/movies", handler.GetMovies)
}

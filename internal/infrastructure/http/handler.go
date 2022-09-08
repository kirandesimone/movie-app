package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

type Handler struct {
	app ports.Api
}

func NewHandler(app ports.Api) *Handler {
	return &Handler{app: app}
}

func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	movies := h.app.GetMovies(ctx)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

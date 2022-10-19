package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

type Handler struct {
	service ports.Api
}

func NewHandler(service ports.Api) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	movies := h.service.GetMovies(ctx)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
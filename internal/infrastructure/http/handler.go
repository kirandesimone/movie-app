package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kirandesimone/mflix/movie-app/internal/ports"
)

type Handler struct {
	service ports.Api
}

func NewHandler(service ports.Api) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMovies(c *fiber.Ctx) error {
	movies := h.service.GetMovies()
	return c.JSON(movies)
}

func (h *Handler) GetTopRatedMovies(c *fiber.Ctx) error {
	topRated := h.service.GetTopRatedMovies()
	return c.JSON(topRated)
}

func (h *Handler) GetDramaMovies(c *fiber.Ctx) error {
	dramaMovies := h.service.GetGenreMovies("Drama")
	return c.JSON(dramaMovies)
}

func (h *Handler) GetActionMovies(c *fiber.Ctx) error {
	actionMovies := h.service.GetGenreMovies("Action")
	return c.JSON(actionMovies)
}

package ports

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	GetMovies(c *fiber.Ctx) error
	GetTopRatedMovies(c *fiber.Ctx) error
	GetDramaMovies(c *fiber.Ctx) error
	GetActionMovies(c *fiber.Ctx) error
}

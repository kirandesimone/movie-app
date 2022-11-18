package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/kirandesimone/mflix/movie-app/internal/infrastructure/db"
	"github.com/kirandesimone/mflix/movie-app/internal/infrastructure/http"
	"github.com/kirandesimone/mflix/movie-app/internal/router"
	"github.com/kirandesimone/mflix/movie-app/internal/service"

	settings "github.com/kirandesimone/mflix/movie-app/internal"
)

// Will possibly create an Application struct for better testing????

func BuildAndRun(config *settings.Cfg) {
	settings.InitLogger()

	store := db.ConnectMongoDb(config.Database.Uri, config.Database.Name)
	service := service.New(store)
	defer store.Disconnect()
	handler := http.NewHandler(service)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.HttpRouter(app, handler)

	app.Listen(fmt.Sprintf(":%d", config.Application.Port))
}

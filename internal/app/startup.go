package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kirandesimone/mflix/movie-app/internal/infrastructure/db"
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

	r := mux.NewRouter()
	r.Handle("/", nil)
	router.HttpRouter(r, service)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Application.Port), r))
}

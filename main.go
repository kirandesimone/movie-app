package main

import (
	settings "github.com/kirandesimone/mflix/movie-app/internal"
	"github.com/kirandesimone/mflix/movie-app/internal/app"
)

var config settings.Cfg

func main() {
	app.BuildAndRun(&config)
}

func init() {
	settings.ReadConfig(&config)
}

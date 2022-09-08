package ports

import "net/http"

type Handler interface {
	GetMovies(w http.ResponseWriter, r *http.Request)
}

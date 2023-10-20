package httphandler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewHttpHandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// docs
	docFS := http.StripPrefix("/docs/", http.FileServer(http.Dir("./docs")))
	r.Get("/docs/*", func(w http.ResponseWriter, r *http.Request) {
		docFS.ServeHTTP(w, r)
	})

	return r
}

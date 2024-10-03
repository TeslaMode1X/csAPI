package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func router() http.Handler {
	r := chi.NewRouter()

	r.Get("/player/{steamID}", playerHandler)

	return r
}

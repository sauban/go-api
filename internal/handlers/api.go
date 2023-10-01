package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/sauban/go-api/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	// Get Account route
	r.Route("/account", func(router chi.Router) {
		// Set up middleware for /account
		router.Use(middleware.Authorization)

		router.Get("/coin", GetCoinBalance)
	})
}

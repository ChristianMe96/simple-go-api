package handlers

import (
	"github.com/ChristianMe96/simple-go-api/internal/middleware"
	"github.com/go-chi/chi/v5"
	chimiddle "github.com/go-chi/chi/v5/middleware"
)

func Handler(router *chi.Mux) {
	router.Use(chimiddle.StripSlashes)

	router.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}

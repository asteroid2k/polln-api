package main

import (
	"github.com/asteroid2k/polln-api/internal/config"
	"github.com/asteroid2k/polln-api/internal/middleware"
	"github.com/asteroid2k/polln-api/internal/polls"
	"github.com/asteroid2k/polln-api/internal/users"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.App) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestLoggerMiddleware)
	mux.Get("/", Home)
	users.RegisterRoutes(app, mux)
	polls.RegisterRoutes(app, mux)
	return mux
}

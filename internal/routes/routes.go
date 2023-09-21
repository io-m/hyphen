package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/io-m/hyphen/internal/features/profiles"
	dependency "github.com/io-m/hyphen/internal/shared/config"
	"github.com/io-m/hyphen/pkg/constants"
)

func ConfigureRoutes(config *dependency.Dependencies) {
	config.GetMux().Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	config.GetMux().Use(middleware.Heartbeat("/ping"))
	config.GetMux().Use(middleware.Logger)
	config.GetMux().Use(middleware.Recoverer)
	config.GetMux().Route(constants.BASE_ROUTE, func(r chi.Router) {
		config.SetRouter(r)
		/* ROUTES COME HERE*/
		profiles.SetAndRunProfileRoutes(config)
	})
}

package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/io-m/hyphen/internal/features/customers"
	"github.com/io-m/hyphen/internal/features/profiles"
	dependency "github.com/io-m/hyphen/internal/shared/dependencies"
	"github.com/io-m/hyphen/pkg/constants"
)

func ConfigureRoutes(dep *dependency.Dependencies) {
	dep.GetMux().Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Accept", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	dep.GetMux().Use(middleware.Heartbeat("/ping"))
	dep.GetMux().Use(middleware.Logger)
	dep.GetMux().Use(middleware.Recoverer)
	dep.GetMux().Route(constants.BASE_ROUTE, func(r chi.Router) {
		dep.SetRouter(r)
		/* ROUTES COME HERE*/
		profiles.SetAndRunProfileRoutes(dep)
		customers.SetAndRunCustomerRoutes(dep)
	})
}

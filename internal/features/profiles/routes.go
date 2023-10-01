package profiles

import (
	"github.com/go-chi/chi/v5"
	profile_handler "github.com/io-m/hyphen/internal/features/profiles/handler"
	dependency "github.com/io-m/hyphen/internal/shared/dependencies"
)

func SetAndRunProfileRoutes(dep *dependency.Dependencies) {
	profileHandler := profile_handler.NewProfileAuthHandler(dep.GetProfileLogic())

	/* AUTH ROUTES */
	dep.GetRouter().Route("/profiles", func(r chi.Router) {
		r.Post("/register", profileHandler.Register)
		r.Post("/login", profileHandler.Login)
		r.Post("/refresh-tokens", profileHandler.RefreshToken)
		r.Put("/password-reset", profileHandler.RefreshToken) // TODO: Implement real handler
		r.Route("/oauth", func(r chi.Router) {
			// r.Use(middlewares.MustAuthenticate(dep.GetProtector()))
			r.Get("/{provider}/login", profileHandler.OAuth)    // TODO: Implement real handler
			r.Get("/{provider}/callback", profileHandler.OAuth) // TODO: Implement real handler
		})

	})
}

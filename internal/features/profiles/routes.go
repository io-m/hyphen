package profiles

import (
	"github.com/go-chi/chi/v5"
	profile_handler "github.com/io-m/hyphen/internal/features/profiles/handler"
	profile_logic "github.com/io-m/hyphen/internal/features/profiles/logic"
	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/config"
)

func SetAndRunProfileRoutes(config *config.AppConfig) {
	profileRepository := profile_repository.NewProfileRepository(config.GetPostgres(), config.GetRedis())
	profileLogic := profile_logic.NewProfileLogic(profileRepository, config.GetTokens(), config.GetProtector())
	profileHandler := profile_handler.NewProfileAuthHandler(profileLogic)

	/* AUTH ROUTES */
	config.GetRouter().Route("/profiles", func(r chi.Router) {
		r.Post("/register", profileHandler.Register)
		r.Post("/login", profileHandler.Login)
		r.Post("/refresh-tokens", profileHandler.RefreshToken)
		r.Put("/password-reset", profileHandler.RefreshToken) // TODO: Implement real handler
		r.Route("/oauth", func(r chi.Router) {
			r.Get("/{provider}/login", profileHandler.OAuth)    // TODO: Implement real handler
			r.Get("/{provider}/callback", profileHandler.OAuth) // TODO: Implement real handler
		})

	})
}

/*
POST /user/register - User Registration
POST /user/login - User Login
POST /user/refresh-token - User Login
POST /user/password-reset - Initiate Password Reset
PUT /user/password-reset/{token} - Complete Password Reset

GET /user/oauth/login - OAuth User Login (e.g. ?provider=google)
GET /user/oauth/callback - OAuth User callback
*/

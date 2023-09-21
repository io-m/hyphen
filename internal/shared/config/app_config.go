package config

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	profile_logic "github.com/io-m/hyphen/internal/features/profiles/logic"
	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/tokens"
)

type AppConfig struct {
	mux          *chi.Mux
	router       chi.Router
	postgres     *sql.DB
	redisClient  *redis.Client
	profileLogic profile_logic.IProfileLogic
}

func NewAppConfig(pg *sql.DB, redis *redis.Client) *AppConfig {
	tokenLogic := tokens.NewTokens(redis)
	protector := tokens.NewProtector()
	return &AppConfig{
		mux:          chi.NewRouter(),
		postgres:     pg,
		redisClient:  redis,
		profileLogic: profile_logic.NewProfileLogic(profile_repository.NewProfileRepository(pg, redis), tokenLogic, protector),
	}
}

func (ac *AppConfig) GetRouter() chi.Router {
	return ac.router
}
func (ac *AppConfig) GetMux() *chi.Mux {
	return ac.mux
}
func (ac *AppConfig) GetPostgres() *sql.DB {
	return ac.postgres
}
func (ac *AppConfig) GetRedis() *redis.Client {
	return ac.redisClient
}
func (ac *AppConfig) GetProfileLogic() profile_logic.IProfileLogic {
	return ac.profileLogic
}
func (ac *AppConfig) SetRouter(router chi.Router) {
	ac.router = router
}

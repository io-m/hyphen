package config

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/tokens"
)

type AppConfig struct {
	mux              *chi.Mux
	router           chi.Router
	protector        tokens.IProtector
	tokens           tokens.ITokens
	postgres         *sql.DB
	redisClient      *redis.Client
	personRepository profile_repository.IProfileRepository
}

func NewAppConfig(pg *sql.DB, redis *redis.Client) *AppConfig {

	return &AppConfig{
		mux:              chi.NewRouter(),
		protector:        tokens.NewProtector(),
		tokens:           tokens.NewTokens(redis),
		postgres:         pg,
		redisClient:      redis,
		personRepository: profile_repository.NewProfileRepository(pg, redis),
	}
}

func (ac *AppConfig) GetRouter() chi.Router {
	return ac.router
}
func (ac *AppConfig) GetMux() *chi.Mux {
	return ac.mux
}
func (ac *AppConfig) GetProtector() tokens.IProtector {
	return ac.protector
}
func (ac *AppConfig) GetTokens() tokens.ITokens {
	return ac.tokens
}
func (ac *AppConfig) GetPostgres() *sql.DB {
	return ac.postgres
}
func (ac *AppConfig) GetRedis() *redis.Client {
	return ac.redisClient
}
func (ac *AppConfig) GetProfileRepository() profile_repository.IProfileRepository {
	return ac.personRepository
}
func (ac *AppConfig) SetRouter(router chi.Router) {
	ac.router = router
}

package dependency

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	profile_logic "github.com/io-m/hyphen/internal/features/profiles/logic"
	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/tokens"
)

type Dependencies struct {
	multiplexer  *chi.Mux
	router       chi.Router
	protector    tokens.IProtector
	postgres     *sql.DB
	redisClient  *redis.Client
	profileLogic profile_logic.IProfileLogic
}

func NewDependencies(pg *sql.DB, redis *redis.Client) *Dependencies {
	tokenLogic := tokens.NewTokens(redis)
	protector := tokens.NewProtector()
	return &Dependencies{
		multiplexer:  chi.NewRouter(),
		postgres:     pg,
		protector:    protector,
		redisClient:  redis,
		profileLogic: profile_logic.NewProfileLogic(profile_repository.NewProfileRepository(pg, redis), tokenLogic, protector),
	}
}

// Getters
func (dep *Dependencies) GetRouter() chi.Router {
	return dep.router
}
func (dep *Dependencies) GetMux() *chi.Mux {
	return dep.multiplexer
}
func (dep *Dependencies) GetProtector() tokens.IProtector {
	return dep.protector
}
func (dep *Dependencies) GetPostgres() *sql.DB {
	return dep.postgres
}
func (dep *Dependencies) GetRedis() *redis.Client {
	return dep.redisClient
}
func (dep *Dependencies) GetProfileLogic() profile_logic.IProfileLogic {
	return dep.profileLogic
}

// Setters
func (dep *Dependencies) SetRouter(router chi.Router) {
	dep.router = router
}

package dependency

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	address_logic "github.com/io-m/hyphen/internal/features/addresses/logic"
	address_repository "github.com/io-m/hyphen/internal/features/addresses/repository"
	customer_logic "github.com/io-m/hyphen/internal/features/customers/logic"
	customer_repository "github.com/io-m/hyphen/internal/features/customers/repository"
	profile_logic "github.com/io-m/hyphen/internal/features/profiles/logic"
	profile_repository "github.com/io-m/hyphen/internal/features/profiles/repository"
	"github.com/io-m/hyphen/internal/shared/tokens"
)

type Dependencies struct {
	multiplexer   *chi.Mux
	router        chi.Router
	protector     tokens.IProtector
	postgres      *sql.DB
	redisClient   *redis.Client
	profileLogic  profile_logic.IProfileLogic
	addressLogic  address_logic.IAddressLogic
	customerLogic customer_logic.ICustomerLogic
}

func NewDependencies(pg *sql.DB, redis *redis.Client) *Dependencies {
	/* SHARED INSTANCES */
	tokenLogic := tokens.NewTokens(redis)
	protector := tokens.NewProtector()
	addressLogic := address_logic.NewAddressLogic(address_repository.NewAddressRepository(pg, redis))
	customerLogic := customer_logic.NewCustomerLogic(customer_repository.NewCustomerRepository(pg, redis), addressLogic)
	/* GLOBAL */
	return &Dependencies{
		multiplexer:   chi.NewRouter(),
		postgres:      pg,
		protector:     protector,
		redisClient:   redis,
		profileLogic:  profile_logic.NewProfileLogic(profile_repository.NewProfileRepository(pg, redis), tokenLogic, protector),
		addressLogic:  addressLogic,
		customerLogic: customerLogic,
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
func (dep *Dependencies) GetAddressLogic() address_logic.IAddressLogic {
	return dep.addressLogic
}
func (dep *Dependencies) GetCustomerLogic() customer_logic.ICustomerLogic {
	return dep.customerLogic
}

// Setters
func (dep *Dependencies) SetRouter(router chi.Router) {
	dep.router = router
}

package customer_repository

import (
	"context"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/io-m/hyphen/internal/shared/models"
)

type customerRepository struct {
	postgres *sql.DB
	redis    *redis.Client
}

func NewCustomerRepository(postgres *sql.DB, redis *redis.Client) *customerRepository {
	return &customerRepository{
		postgres: postgres,
		redis:    redis,
	}
}

type ICustomerRepository interface {
	FindAllCustomers(ctx context.Context) ([]models.Customer, error)
	FindCustomerById(ctx context.Context, customerId int) (models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.Customer) (int, error)
	UpdateCustomerById(ctx context.Context, customerId int, customer *models.Customer) (models.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId int) (bool, error)
}

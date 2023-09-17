package customer_repository

import (
	"context"

	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/jmoiron/sqlx"
)

type postgresCustomerRepository struct {
	postgres *sqlx.DB
	// redis    *redis.Client
}

type ICustomerRepository interface {
	FindAllCustomers(ctx context.Context) ([]models.Customer, error)
	FindCustomerById(ctx context.Context, customerId uint) (models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.Customer) (models.Customer, error)
	UpdateCustomerById(ctx context.Context, customerId uint, customer *models.Customer) (models.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId uint) (bool, error)
}

package customer_repository

import (
	"context"

	"github.com/io-m/hyphen/internal/shared/models"
)

func (cr *postgresCustomerRepository) FindAllCustomers(ctx context.Context) ([]models.Customer, error) {
	return []models.Customer{models.EmptyCustomer()}, nil
}
func (cr *postgresCustomerRepository) FindCustomerById(ctx context.Context, customerId uint) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cr *postgresCustomerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cr *postgresCustomerRepository) UpdateCustomerById(ctx context.Context, customerId uint, customer *models.Customer) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cr *postgresCustomerRepository) DeleteCustomerById(ctx context.Context, customerId uint) (bool, error) {
	return false, nil
}

package customer_repository

import (
	"context"
	_ "embed"
	"log/slog"

	"github.com/io-m/hyphen/internal/shared/models"
)

//go:embed queries/create_customer.psql
var CREATE_CUSTOMER string

//go:embed queries/find_all_customers.psql
var FIND_ALL_CUSTOMERS string

func (cr *customerRepository) FindAllCustomers(ctx context.Context) ([]models.Customer, error) {

	rows, err := cr.postgres.Query(FIND_ALL_CUSTOMERS)
	if err != nil {
		slog.Debug("Failed to retrieve customers: %v", err)
		return nil, err
	}
	defer rows.Close()
	var customers []models.Customer
	for rows.Next() {
		var customer models.Customer
		if err := rows.Scan(&customer.ProfileID, &customer.AddressId, &customer.FullName, &customer.Gender, &customer.Bio, &customer.Blacklisted); err != nil {
			slog.Debug("Failed to scan customer: %v", err)
			return nil, err
		}
		customers = append(customers, customer)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		slog.Debug("Error iterating rows: %v", err)
		return nil, err
	}

	return customers, nil
}
func (cr *customerRepository) FindCustomerById(ctx context.Context, customerId int) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cr *customerRepository) CreateCustomer(ctx context.Context, customer *models.Customer) (int, error) {
	var id int
	err := cr.postgres.QueryRow(CREATE_CUSTOMER, customer.ProfileID, customer.AddressId, customer.FullName, *customer.Gender, *customer.Bio, customer.Blacklisted).Scan(&id)
	if err != nil {
		slog.Debug("Failed to insert customer: %v", err)
		return -1, err
	}

	return id, nil
}
func (cr *customerRepository) UpdateCustomerById(ctx context.Context, customerId int, customer *models.Customer) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cr *customerRepository) DeleteCustomerById(ctx context.Context, customerId int) (bool, error) {
	return false, nil
}

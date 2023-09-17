package customer_logic

import (
	"context"

	customer_repository "github.com/io-m/hyphen/internal/features/customers/repository"
	"github.com/io-m/hyphen/internal/shared/models"
)

type ICustomerLogic interface {
	GetAllCustomers(ctx context.Context) ([]models.Customer, error)
	GetCustomerById(ctx context.Context, customerId uint) (models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.Customer) (models.Customer, error)
	UpdateCustomerById(ctx context.Context, customerId uint, customer *models.Customer) (models.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId uint) (bool, error)

	GetCustomersBooking(ctx context.Context, customerId uint) (models.Booking, error)
}

type customerLogic struct {
	customerRepository customer_repository.ICustomerRepository
	bookingRepository  any // TODO: implement booking repository
}

func NewUserLogic(customerRepository customer_repository.ICustomerRepository, bookingRepository any) *customerLogic {
	return &customerLogic{
		customerRepository: customerRepository,
		bookingRepository:  bookingRepository,
	}
}

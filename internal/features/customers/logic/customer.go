package customer_logic

import (
	"context"

	address_logic "github.com/io-m/hyphen/internal/features/addresses/logic"
	customer_repository "github.com/io-m/hyphen/internal/features/customers/repository"
	"github.com/io-m/hyphen/internal/shared"
	"github.com/io-m/hyphen/internal/shared/models"
)

type ICustomerLogic interface {
	GetAllCustomers(ctx context.Context) ([]models.Customer, error)
	GetCustomerById(ctx context.Context, customerId uint) (models.Customer, error)
	CreateCustomer(ctx context.Context, customer *models.CustomerRequest) (shared.ItemID, error)
	UpdateCustomerById(ctx context.Context, customerId uint, customer *models.Customer) (models.Customer, error)
	DeleteCustomerById(ctx context.Context, customerId uint) (bool, error)

	// GetCustomersBooking(ctx context.Context, customerId uint) (models.Booking, error)
}

type customerLogic struct {
	customerRepository customer_repository.ICustomerRepository
	addressLogic       address_logic.IAddressLogic
	// bookingRepository  any // TODO: implement booking repository
}

func NewCustomerLogic(customerRepository customer_repository.ICustomerRepository, addressLogic address_logic.IAddressLogic) *customerLogic {
	return &customerLogic{
		customerRepository: customerRepository,
		addressLogic:       addressLogic,
		// bookingRepository:  bookingRepository,
	}
}

func (cl *customerLogic) GetAllCustomers(ctx context.Context) ([]models.Customer, error) {
	return cl.customerRepository.FindAllCustomers(ctx)
}
func (cl *customerLogic) GetCustomerById(ctx context.Context, customerId uint) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cl *customerLogic) CreateCustomer(ctx context.Context, customerRequest *models.CustomerRequest) (shared.ItemID, error) {
	// Make some validations on inputs
	item, err := cl.addressLogic.GenerateAddress(ctx, &customerRequest.Address)
	if err != nil {
		return shared.ItemID{ID: -1}, err
	}
	customer := models.MapCustomerRequestToCustomerModel(*customerRequest, item.ID)
	customer.Blacklisted = false
	customerItemId, err := cl.customerRepository.CreateCustomer(ctx, &customer)
	if err != nil {
		return shared.ItemID{ID: customerItemId}, err
	}
	return shared.ItemID{ID: customerItemId}, nil
}
func (cl *customerLogic) UpdateCustomerById(ctx context.Context, customerId uint, customer *models.Customer) (models.Customer, error) {
	return models.EmptyCustomer(), nil
}
func (cl *customerLogic) DeleteCustomerById(ctx context.Context, customerId uint) (bool, error) {
	return false, nil
}

// GetCustomersBooking(ctx context.Context, customerId uint) (models.Booking, error)

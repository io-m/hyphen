package customer_handler

import (
	"fmt"
	"net/http"

	customer_logic "github.com/io-m/hyphen/internal/features/customers/logic"
	"github.com/io-m/hyphen/internal/shared/models"
	"github.com/io-m/hyphen/pkg/helpers"
)

type ICustomerHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	UpdateById(w http.ResponseWriter, r *http.Request)
	DeleteById(w http.ResponseWriter, r *http.Request)
	GetBookings(w http.ResponseWriter, r *http.Request)
}

type customerHandler struct {
	customerLogic customer_logic.ICustomerLogic
}

func NewCustomerHandler(customerLogic customer_logic.ICustomerLogic) *customerHandler {
	return &customerHandler{
		customerLogic: customerLogic,
	}
}

func (ch *customerHandler) Create(w http.ResponseWriter, r *http.Request) {
	customerRequest, err := helpers.DecodePayload[*models.CustomerRequest](w, r)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("error while decoding payload: %w", err), http.StatusBadRequest)
		return
	}
	createdCustomerId, err := ch.customerLogic.CreateCustomer(r.Context(), customerRequest)
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not create customer: %w", err), http.StatusInternalServerError)
		return
	}
	helpers.SuccessResponse(w, createdCustomerId, "customer successfully created", http.StatusCreated)
}
func (ch *customerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.customerLogic.GetAllCustomers(r.Context())
	if err != nil {
		helpers.ErrorResponse(w, fmt.Errorf("could not find customers: %w", err), http.StatusNotFound)
		return
	}
	helpers.SuccessResponse(w, customers, "all customers found", http.StatusOK)
}
func (ch *customerHandler) GetById(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) UpdateById(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) DeleteById(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) GetBookings(w http.ResponseWriter, r *http.Request) {

}

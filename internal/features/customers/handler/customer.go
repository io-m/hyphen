package customer_handler

import (
	"net/http"

	customer_logic "github.com/io-m/hyphen/internal/features/customers/logic"
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

func NewUserAuthHandler(customerLogic customer_logic.ICustomerLogic) *customerHandler {
	return &customerHandler{
		customerLogic: customerLogic,
	}
}

func (ch *customerHandler) Create(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) GetById(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) UpdateById(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) DeleteById(w http.ResponseWriter, r *http.Request) {

}
func (ch *customerHandler) GetBookings(w http.ResponseWriter, r *http.Request) {

}

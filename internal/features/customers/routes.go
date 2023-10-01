package customers

import (
	"github.com/go-chi/chi/v5"
	customer_handler "github.com/io-m/hyphen/internal/features/customers/handler"
	dependency "github.com/io-m/hyphen/internal/shared/dependencies"
)

func SetAndRunCustomerRoutes(dep *dependency.Dependencies) {
	customerHandler := customer_handler.NewCustomerHandler(dep.GetCustomerLogic())

	/* CUSTOMER ROUTES */
	dep.GetRouter().Route("/customers", func(r chi.Router) {
		r.Post("/", customerHandler.Create)
		r.Get("/", customerHandler.GetAll)
		r.Get("/{id}", customerHandler.GetById)
		r.Put("/{id}", customerHandler.UpdateById)
		r.Delete("/{id}", customerHandler.DeleteById)
	})
}

/*
POST /customers - Register a new customer.
GET /customers/{customerId} - Get the profile of a specific customer.
PUT /customers/{customerId} - Update the profile of a specific customer.
DELETE /customers/{customerId} - Delete the profile of a specific customer.

GET /customers/{customerId}/bookings - List All Bookings for a Customer
*/

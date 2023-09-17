package models

import "time"

type Customer struct {
	ID          int        `json:"id,omitempty" db:"id"`
	UserID      uint       `json:"user_id,omitempty" db:"user_id"`
	FirstName   string     `json:"first_name,omitempty" db:"first_name"`
	LastName    string     `json:"last_name,omitempty" db:"last_name"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	BirthDate   *time.Time `json:"birth_date,omitempty" db:"birth_date"`
}

func EmptyCustomer() Customer {
	return Customer{}
}

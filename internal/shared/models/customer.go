package models

import "time"

type Customer struct {
	ProfileID   uint       `json:"profile_id,omitempty" db:"profile_id"`
	FirstName   string     `json:"first_name,omitempty" db:"first_name"`
	LastName    string     `json:"last_name,omitempty" db:"last_name"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	BirthDate   *time.Time `json:"birth_date,omitempty" db:"birth_date"`
	Blacklisted bool       `json:"blacklisted,omitempty" db:"blacklisted"`
}

type CustomerUpdateRequest struct {
	FirstName   *string    `json:"first_name,omitempty" db:"first_name"`
	LastName    *string    `json:"last_name,omitempty" db:"last_name"`
	PhoneNumber *string    `json:"phone_number,omitempty" db:"phone_number"`
	Bio         *string    `json:"bio,omitempty" db:"bio"`
	BirthDate   *time.Time `json:"birth_date,omitempty" db:"birth_date"`
	UpdatedAt   *string    `json:"updated_at,omitempty"`
}

func EmptyCustomer() Customer {
	return Customer{}
}

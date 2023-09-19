package models

import "time"

type Review struct {
	ID         int        `json:"id,omitempty" db:"id"`
	CustomerID int        `json:"customer_id,omitempty" db:"customer_id"`
	Rating     int        `json:"rating,omitempty" db:"rating"`
	Comment    *string    `json:"comment,omitempty" db:"comment"`
	BusinessID *int       `json:"business_id,omitempty" db:"business_id"`
	EmployeeID *int       `json:"employee_id,omitempty" db:"employee_id"`
	CreatedAt  *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

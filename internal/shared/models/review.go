package models

import "time"

type Review struct {
	ID         int        `json:"id" db:"id"`
	CustomerID int        `json:"customer_id" db:"customer_id"`
	BusinessID int        `json:"business_id" db:"business_id"`
	Rating     int        `json:"rating" db:"rating"`
	Comment    *string    `json:"comment" db:"comment"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}

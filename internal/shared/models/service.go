package models

import "time"

type Service struct {
	ID          int           `json:"id,omitempty" db:"id"`
	BusinessID  int           `json:"business_id,omitempty" db:"business_id"`
	CategoryID  int           `json:"category_id,omitempty" db:"category_id"`
	Name        string        `json:"name,omitempty" db:"name"`
	Description string        `json:"description,omitempty" db:"description"`
	Price       float64       `json:"price,omitempty" db:"price"`
	Duration    time.Duration `json:"duration,omitempty" db:"duration"` // Might need custom type
}

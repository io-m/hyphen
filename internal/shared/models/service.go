package models

import "time"

type Service struct {
	ID          int        `json:"id,omitempty" db:"id"`
	BusinessID  int        `json:"business_id,omitempty" db:"business_id"` // this is profile id
	CategoryID  int        `json:"category_id,omitempty" db:"category_id"`
	Name        string     `json:"name,omitempty" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	Price       float64    `json:"price,omitempty" db:"price"`
	Duration    int        `json:"duration,omitempty" db:"duration"`
	CreatedAt   *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

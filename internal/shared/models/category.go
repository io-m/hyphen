package models

import "time"

type Category struct {
	ID        int        `json:"id,omitempty" db:"category_id"`
	Name      string     `json:"name,omitempty" db:"category_name"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

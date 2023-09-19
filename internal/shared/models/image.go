package models

import "time"

type Image struct {
	ID        int        `json:"id,omitempty" db:"id"`
	URL       string     `json:"url,omitempty" db:"url"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

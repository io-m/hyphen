package models

import "time"

type Image struct {
	ID        int        `json:"id" db:"id"`
	URL       string     `json:"url" db:"url"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}

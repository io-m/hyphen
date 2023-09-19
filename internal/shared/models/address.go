package models

import "time"

type Address struct {
	ID        int        `json:"id,omitempty" db:"id"`
	Street    string     `json:"street,omitempty" db:"street"`
	City      string     `json:"city,omitempty" db:"city"`
	State     string     `json:"state,omitempty" db:"state"`
	ZipCode   string     `json:"zip_code,omitempty" db:"zip_code"`
	Country   string     `json:"country,omitempty" db:"country"`
	Latitude  float64    `json:"latitude,omitempty" db:"latitude"`
	Longitude float64    `json:"longitude,omitempty" db:"longitude"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

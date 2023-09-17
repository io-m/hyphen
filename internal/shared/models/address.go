package models

type Address struct {
	ID        int     `json:"id" db:"id"`
	Street    string  `json:"street" db:"street"`
	City      string  `json:"city" db:"city"`
	State     string  `json:"state" db:"state"`
	ZipCode   string  `json:"zip_code" db:"zip_code"`
	Country   string  `json:"country" db:"country"`
	Latitude  float64 `json:"latitude" db:"latitude"`
	Longitude float64 `json:"longitude" db:"longitude"`
}

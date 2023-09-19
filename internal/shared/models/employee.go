package models

type Employee struct {
	ID         uint    `json:"id,omitempty" db:"id"` //Foreign key to Users but also primary key
	FirstName  string  `json:"first_name,omitempty" db:"first_name"`
	LastName   string  `json:"last_name,omitempty" db:"last_name"`
	BusinessID int     `json:"business_id,omitempty" db:"business_id"`
	Experience float32 `json:"experience,omitempty" db:"experience"`
}

package models

type Employee struct {
	ID         int     `json:"id,omitempty" db:"id"`
	UserID     uint    `json:"user_id,omitempty" db:"user_id"`
	FirstName  string  `json:"first_name,omitempty" db:"first_name"`
	LastName   string  `json:"last_name,omitempty" db:"last_name"`
	BusinessID int     `json:"business_id,omitempty" db:"business_id"`
	Experience float32 `json:"experience,omitempty" db:"experience"`
}

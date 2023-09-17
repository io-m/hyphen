package models

type Employee struct {
	ID         int     `json:"id" db:"id"`
	ProfileID  int     `json:"profile_id" db:"profile_id"`
	BusinessID int     `json:"business_id" db:"business_id"`
	Experience float32 `json:"experience" db:"experience"`
}

package models

import "time"

type Customer struct {
	ID          int        `json:"_id" db:"_id"`
	ProfileID   int        `json:"profile_id" db:"profile_id"`
	PhoneNumber *string    `json:"phone_number" db:"phone_number"`
	Bio         *string    `json:"_bio" db:"_bio"`
	BirthDate   *time.Time `json:"_birth_date" db:"_birth_date"`
}

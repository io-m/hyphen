package models

import "time"

type Profile struct {
	ID        int64      `json:"id,omitempty" db:"id"`
	Email     string     `json:"email,omitempty" db:"email"`
	Password  string     `json:"password,omitempty" db:"hashed_password"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type ProfileUpdateRequest struct {
	Email       *string `json:"email,omitempty"`
	OldPassword *string `json:"old_password,omitempty"`
	NewPassword *string `json:"new_password,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

func EmptyProfile() Profile {
	return Profile{}
}

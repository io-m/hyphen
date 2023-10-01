package models

// Gender is a custom type for representing gender using constants.
type Gender int

// Define constants for each gender.
const (
	Male Gender = iota
	Female
)

func (g Gender) String() string {
	// Create a string representation for each gender.
	names := [...]string{"Male", "Female"}
	if g < Male || g > Female {
		return "Unknown"
	}
	return names[g]
}

type Customer struct {
	ProfileID   uint    `json:"profile_id,omitempty" db:"profile_id"`
	FullName    string  `json:"full_name,omitempty" db:"full_name"`
	Gender      *Gender `json:"gender,omitempty" db:"gender"`
	Bio         *string `json:"bio,omitempty" db:"bio"`
	Blacklisted bool    `json:"blacklisted,omitempty" db:"blacklisted"`
}

type CustomerUpdateRequest struct {
	FullName  *string `json:"full_name,omitempty" db:"full_name"`
	Bio       *string `json:"bio,omitempty" db:"bio"`
	Gender    *Gender `json:"gender,omitempty" db:"gender"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func EmptyCustomer() Customer {
	return Customer{}
}

package models

type Customer struct {
	ProfileID   uint    `json:"profile_id,omitempty" db:"profile_id"`
	FullName    string  `json:"full_name,omitempty" db:"full_name"`
	Bio         *string `json:"bio,omitempty" db:"bio"`
	Blacklisted bool    `json:"blacklisted,omitempty" db:"blacklisted"`
}

type CustomerUpdateRequest struct {
	FullName  *string `json:"full_name,omitempty" db:"full_name"`
	Bio       *string `json:"bio,omitempty" db:"bio"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func EmptyCustomer() Customer {
	return Customer{}
}

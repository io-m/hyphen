package models

type Customer struct {
	ProfileID   int     `json:"profile_id,omitempty" db:"profile_id"`
	FullName    string  `json:"full_name,omitempty" db:"full_name"`
	Gender      *string `json:"gender,omitempty" db:"gender"`
	Bio         *string `json:"bio,omitempty" db:"bio"`
	Blacklisted bool    `json:"blacklisted,omitempty" db:"blacklisted"`
	AddressId   int     `json:"address_id,omitempty" db:"address_id"`
}

type CustomerRequest struct {
	ProfileID   int            `json:"profile_id,omitempty" db:"profile_id"`
	FullName    string         `json:"full_name,omitempty" db:"full_name"`
	Gender      *string        `json:"gender,omitempty" db:"gender"`
	Bio         *string        `json:"bio,omitempty" db:"bio"`
	Blacklisted bool           `json:"blacklisted,omitempty" db:"blacklisted"`
	Address     AddressRequest `json:"address"`
}

type CustomerUpdateRequest struct {
	FullName  *string `json:"full_name,omitempty" db:"full_name"`
	Bio       *string `json:"bio,omitempty" db:"bio"`
	Gender    *string `json:"gender,omitempty" db:"gender"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func MapCustomerRequestToCustomerModel(customerRequest CustomerRequest, addressId int) Customer {
	return Customer{
		ProfileID:   customerRequest.ProfileID,
		FullName:    customerRequest.FullName,
		Gender:      customerRequest.Gender,
		Bio:         customerRequest.Bio,
		Blacklisted: customerRequest.Blacklisted,
		AddressId:   addressId,
	}
}

func EmptyCustomer() Customer {
	return Customer{}
}

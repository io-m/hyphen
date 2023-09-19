package models

type Business struct {
	ProfileID      uint     `json:"profile_id,omitempty" db:"profile_id"`
	AddressID      uint     `json:"address_id,omitempty" db:"address_id"`
	Description    string   `json:"description,omitempty" db:"description"`
	OperatingHours string   `json:"operating_hours,omitempty" db:"operating_hours"`
	RatingAvg      float64  `json:"rating_avg,omitempty" db:"rating_avg"`
	PaymentOptions []string `json:"payment_options,omitempty" db:"payment_options"`
	Blacklisted    bool     `json:"blacklisted,omitempty" db:"blacklisted"`
}

type BusinessUpdateRequest struct {
	Description    *string   `json:"description,omitempty" db:"description"`
	OperatingHours *string   `json:"operating_hours,omitempty" db:"operating_hours"`
	RatingAvg      *float64  `json:"rating_avg,omitempty" db:"rating_avg"`
	PaymentOptions []*string `json:"payment_options,omitempty" db:"payment_options"`
	UpdatedAt      *string   `json:"updated_at,omitempty"`
}

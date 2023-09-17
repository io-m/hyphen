package models

type Business struct {
	ID             uint     `json:"id,omitempty" db:"id"`
	UserID         uint     `json:"user_id,omitempty" db:"user_id"`
	AddressID      uint     `json:"address_id,omitempty" db:"address_id"`
	Description    string   `json:"description,omitempty" db:"description"`
	OperatingHours string   `json:"operating_hours,omitempty" db:"operating_hours"`
	RatingAvg      float64  `json:"rating_avg,omitempty" db:"rating_avg"`
	PaymentOptions []string `json:"payment_options,omitempty" db:"payment_options"`
}

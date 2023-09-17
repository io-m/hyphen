package models

type Business struct {
	ID             int      `json:"id" db:"id"`
	ProfileID      int      `json:"profile_id" db:"profile_id"`
	Description    string   `json:"description" db:"description"`
	OperatingHours string   `json:"operating_hours" db:"operating_hours"`
	RatingAvg      float64  `json:"rating_avg" db:"rating_avg"`
	PaymentOptions []string `json:"payment_options" db:"payment_options"`
}

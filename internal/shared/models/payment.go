package models

import "time"

type Payment struct {
	ID            int        `json:"id,omitempty" db:"id"`
	BookingID     int        `json:"booking_id,omitempty" db:"booking_id"`
	Amount        float64    `json:"amount,omitempty" db:"amount"`
	Status        string     `json:"status,omitempty" db:"status"`
	PaymentMethod string     `json:"payment_method,omitempty" db:"payment_method"`
	TransactionID string     `json:"transaction_id,omitempty" db:"transaction_id"`
	CreatedAt     *time.Time `json:"created_at,omitempty" db:"created_at"`
}

package models

import "time"

type Payment struct {
	ID            int       `json:"id" db:"id"`
	BookingID     int       `json:"booking_id" db:"booking_id"`
	Amount        float64   `json:"amount" db:"amount"`
	Status        string    `json:"status" db:"status"`
	PaymentMethod string    `json:"payment_method" db:"payment_method"`
	TransactionID string    `json:"transaction_id" db:"transaction_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

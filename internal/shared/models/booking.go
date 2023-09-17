package models

import "time"

type Booking struct {
	ID         int        `json:"_id" db:"_id"`
	UserID     int        `json:"user_id" db:"user_id"`
	ServiceID  int        `json:"service_id" db:"service_id"`
	EmployeeID int        `json:"employee_id" db:"employee_id"`
	DateTime   time.Time  `json:"date_time" db:"date_time"`
	Status     string     `json:"status" db:"status"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at"`
}

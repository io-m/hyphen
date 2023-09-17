package models

import "time"

type Booking struct {
	ID         int        `json:"_id,omitempty" db:"_id"`
	UserID     int        `json:"user_id,omitempty" db:"user_id"`
	ServiceID  int        `json:"service_id,omitempty" db:"service_id"`
	EmployeeID int        `json:"employee_id,omitempty" db:"employee_id"`
	DateTime   time.Time  `json:"date_time,omitempty" db:"date_time"`
	Status     string     `json:"status" db,omitempty:"status"`
	CreatedAt  time.Time  `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

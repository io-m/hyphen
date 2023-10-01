package models

type CustomerServiceWishlist struct {
	CustomerId int `json:"customer_id,omitempty" db:"customer_id"`
	ServiceId  int `json:"service_id,omitempty" db:"service_id"`
}

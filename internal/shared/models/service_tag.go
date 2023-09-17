package models

type ServiceTag struct {
	ServiceID int `json:"service_id" db:"service_id"`
	TagID     int `json:"tag_id" db:"tag_id"`
}

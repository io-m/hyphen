package models

type ServiceImage struct {
	ServiceID int `json:"service_id" db:"service_id"`
	ImageID   int `json:"image_id" db:"image_id"`
}

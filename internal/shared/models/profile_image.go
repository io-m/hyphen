package models

type ProfileImage struct {
	ProfileID int `json:"profile_id" db:"profile_id"`
	ImageID   int `json:"image_id" db:"image_id"`
}

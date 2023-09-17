package models

type UserImage struct {
	UserID  int `json:"user_id" db:"user_id"`
	ImageID int `json:"image_id" db:"image_id"`
}

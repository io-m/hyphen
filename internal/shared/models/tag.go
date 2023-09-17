package models

type Tag struct {
	TagID   int    `json:"tag_id" db:"tag_id"`
	TagName string `json:"tag_name" db:"tag_name"`
}

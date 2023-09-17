package models

type Tag struct {
	TagID   int    `json:"tag_id,omitempty" db:"tag_id"`
	TagName string `json:"tag_name,omitempty" db:"tag_name"`
}

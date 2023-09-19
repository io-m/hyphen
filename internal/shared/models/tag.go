package models

import "time"

type Tag struct {
	TagID     int        `json:"tag_id,omitempty" db:"tag_id"`
	TagName   string     `json:"tag_name,omitempty" db:"tag_name"`
	CreatedAt *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

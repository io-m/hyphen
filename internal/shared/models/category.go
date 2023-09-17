package models

type Category struct {
	ID   int    `json:"id" db:"category_id"`
	Name string `json:"name" db:"category_name"`
}

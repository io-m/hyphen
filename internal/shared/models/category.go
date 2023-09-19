package models

type Category struct {
	ID   int    `json:"id,omitempty" db:"category_id"`
	Name string `json:"name,omitempty" db:"category_name"`
	// Icon string `json:"icon,omitempty"`
}

package models

type Service struct {
	ID          int      `json:"id" db:"id"`
	BusinessID  int      `json:"business_id" db:"business_id"`
	CategoryID  int      `json:"category_id" db:"category_id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Price       float64  `json:"price" db:"price"`
	Duration    string   `json:"duration" db:"duration"` // Might need custom type
	PhotoURLs   []string `json:"photo_urls" db:"photo_urls"`
}

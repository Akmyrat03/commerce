package model

type Product struct {
	ID           int     `json:"id" db:"id"`
	Name         string  `json:"name" db:"name"`
	Description  string  `json:"description" db:"description"`
	Image        string  `json:"image" db:"image"`
	Price        float64 `json:"price" db:"price"`
	CategoryID   int     `json:"-" db:"category_id"`
	CategoryName string  `json:"category_name" db:"category_name"`
}

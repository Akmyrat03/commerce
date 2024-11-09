package model

type Product struct {
	ID          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image"`
	Price       int    `json:"price" db:"price"`
	CategoryID  int    `json:"category_id" db:"category_id"`
}

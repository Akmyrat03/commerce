package model

import "time"

type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" binding:"required" db:"name"`
	Image     string    `json:"image" db:"image"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

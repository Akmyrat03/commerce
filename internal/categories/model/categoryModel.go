package model

import "time"

type Category struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" binding:"required" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

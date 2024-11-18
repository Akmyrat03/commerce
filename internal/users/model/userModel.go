package model

import "time"

type User struct {
	ID          int       `json:"id" db:"id"`
	Username    string    `json:"username" binding:"required" db:"username"`
	PhoneNumber string    `json:"phone_number" db:"phone_number"`
	Password    string    `json:"password" binding:"required" db:"password"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Role        string    `json:"role" db:"role"`
}

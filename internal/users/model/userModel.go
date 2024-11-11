// package model

// import "time"

// type UserDTO struct {
// 	ID        int       `json:"id"`
// 	Username  string    `json:"username"`
// 	Email     string    `json:"email"`
// 	Password  string    `json:"password"`
// 	CreatedAt time.Time `json:"created_at"`
// 	Role      string    `json:"role"`
// }

// type UserDAO struct {
// 	ID        int       `db:"id"`
// 	Username  string    `db:"username"`
// 	Email     string    `db:"email"`
// 	Password  string    `db:"password"`
// 	CreatedAt time.Time `db:"created_at"`
// 	Role      string    `db:"role"`
// }

// // ToServer is
// func (d *UserDAO) ToServer() *UserDTO {
// 	return &UserDTO{
// 		ID:        d.ID,
// 		Username:  d.Username,
// 		Email:     d.Email,
// 		Password:  d.Password,
// 		CreatedAt: d.CreatedAt,
// 	}
// }

package model

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	Username  string    `json:"username" binding:"required" db:"username"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" binding:"required" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Role      string    `json:"role" db:"role"`
}

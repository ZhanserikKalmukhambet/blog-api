package model

type User struct {
	ID        int64  `json:"id" db:"id"`
	Username  string `json:"user_name" db:"user_name" binding:"required"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	Password  string `json:"password" db:"password" binding:"required"`
}

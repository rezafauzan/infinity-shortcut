package dto

import "time"

type CreateUserDTO struct {
	FirstName       string `json:"first_name" db:"first_name" binding:"first_name"`
	LastName        string `json:"last_name" db:"last_name" binding:"last_name"`
	Email           string `json:"email" db:"email" binding:"email"`
	Password        string `json:"password" db:"password" binding:"password"`
	PasswordConfirm string `json:"password_confirm" db:"password_confirm" binding:"password_confirm"`
}

type CreateUserResponseDTO struct {
	Id        int       `json:"id" db:"id"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
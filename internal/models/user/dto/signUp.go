package dto

import "time"

type SignUpRequest struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required"`
	Birthdate time.Time `json:"birthdate" validate:"required"`
	Gender    string    `json:"gender" validate:"oneof=male female prefer_not_to"`
}

type SignUpResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

package auth

import "github.com/golang-jwt/jwt/v5"

type UserJwtAuth struct {
	UserID int64  `json:"userID"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

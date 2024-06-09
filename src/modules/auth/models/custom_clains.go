package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims define las reclamaciones personalizadas que incluyen los campos adicionales del modelo User
type CustomClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

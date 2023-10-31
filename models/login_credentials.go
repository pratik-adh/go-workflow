package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type LoginCredentials struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password"` 
}

type AuthCustomClaims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ID		uuid.UUID 
	jwt.StandardClaims
}
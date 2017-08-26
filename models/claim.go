package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim estructura de peticiones
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}

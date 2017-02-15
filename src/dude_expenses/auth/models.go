package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Value string `json:"token,omitempty"`
}

type AuthClaims struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

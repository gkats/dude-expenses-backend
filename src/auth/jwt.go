package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

// TODO add this to a secrets file. Get it from env in production
var authSecret = "my-secret-key"

type JwtError struct {
	Err error
}

func (e JwtError) Error() string {
	return e.Err.Error()
}

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (t *TokenService) CreateToken() (string, error) {
	claims := jwt.MapClaims{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(authSecret))
}

func (t *TokenService) ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid JWT signing method")
		}
		return []byte(authSecret), nil
	})

	if err != nil || !token.Valid {
		return &JwtError{Err: errors.New("Invalid JWT token")}
	}
	return nil
}

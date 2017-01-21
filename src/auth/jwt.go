package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

// TODO add this to a secrets file. Get it from env in production
// or pass it to NewTokenService one level up
var authSecret = "my-secret-key"

type JwtError struct {
	Err error
}

func (e JwtError) Error() string {
	return e.Err.Error()
}

type TokenService struct {
	claims AuthClaims
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (t *TokenService) GetClaims() AuthClaims {
	return t.claims
}

func (t *TokenService) SetClaims(claims *AuthClaims) {
	t.claims = *claims
}

func (t *TokenService) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t.claims)
	return token.SignedString([]byte(authSecret))
}

func (t *TokenService) ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, parser)
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		t.SetClaims(claims)
	} else {
		return &JwtError{Err: errors.New("Invalid JWT token")}
	}
	return nil
}

func parser(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("Invalid JWT signing method")
	}
	return []byte(authSecret), nil
}

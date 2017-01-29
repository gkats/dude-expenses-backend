package auth

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
)

type JwtError struct {
	Err error
}

func (e JwtError) Error() string {
	return e.Err.Error()
}

type tokenService struct {
	secret string
	claims AuthClaims
}

func newTokenService(secret string) *tokenService {
	return &tokenService{secret: secret}
}

func (t *tokenService) GetClaims() AuthClaims {
	return t.claims
}

func (t *tokenService) SetClaims(claims *AuthClaims) {
	t.claims = *claims
}

func (t *tokenService) CreateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t.claims)
	return token.SignedString([]byte(t.secret))
}

func (t *tokenService) ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &AuthClaims{}, parserFor(t.secret))
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

func parserFor(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid JWT signing method")
		}
		return []byte(secret), nil
	}
}

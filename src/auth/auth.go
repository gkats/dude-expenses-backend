package auth

import (
	"app"
	"errors"
	"net/http"
	"regexp"
)

func GenerateToken(id string) (*Token, error) {
	tokenService := NewTokenService()
	tokenService.SetClaims(&AuthClaims{UserId: id})
	tokenString, err := tokenService.CreateToken()
	if err != nil {
		return nil, err
	}
	return &Token{Value: tokenString}, nil
}

func ParseTokenUserId(token string) (string, error) {
	tokenService := NewTokenService()
	if err := tokenService.ParseToken(token); err != nil {
		return "", err
	}
	return tokenService.GetClaims().UserId, nil
}

// TODO add this as middleware
func Authenticate(env *app.Env, r *http.Request) error {
	authorization := r.Header.Get("Authorization")
	regex, _ := regexp.Compile("^Bearer (.+)")
	matches := regex.FindStringSubmatch(authorization)
	if len(matches) < 1 {
		return errors.New("ERROR")
	}

	// TODO provide a public method in auth...
	// also hide tokenService from the outside world
	userId, err := ParseTokenUserId(matches[1])
	if err != nil {
		return errors.New("ERROR")
	}
	env.SetUserId(userId)
	return nil
}

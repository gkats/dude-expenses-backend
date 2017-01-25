package auth

import (
	"dude_expenses/app"
	"net/http"
	"regexp"
)

func GenerateToken(id string) (*Token, error) {
	tokenService := newTokenService()
	tokenService.SetClaims(&AuthClaims{UserId: id})
	tokenString, err := tokenService.CreateToken()
	if err != nil {
		return nil, err
	}
	return &Token{Value: tokenString}, nil
}

func ParseTokenUserId(token string) (string, error) {
	tokenService := newTokenService()
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
		return &AuthError{msg: "Invalid Authorization header"}
	}

	userId, err := ParseTokenUserId(matches[1])
	if err != nil {
		return &AuthError{msg: "Invalid authorization token"}
	}
	env.SetUserId(userId)
	return nil
}

type AuthError struct {
	msg string
}

func (e *AuthError) Error() string {
	return e.msg
}

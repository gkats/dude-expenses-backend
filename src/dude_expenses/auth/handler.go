package auth

import (
	"dude_expenses/app"
	"net/http"
	"regexp"
)

type authHandler struct {
	env  *app.Env
	next app.Handler
}

func (h authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) app.Response {
	authorization := r.Header.Get("Authorization")
	regex, _ := regexp.Compile("^Bearer (.+)")
	matches := regex.FindStringSubmatch(authorization)
	if len(matches) < 1 {
		return app.Unauthorized()
	}

	userId, err := parseTokenUserId(matches[1], h.env.GetAuthSecret())
	if err != nil {
		return app.Unauthorized()
	}
	h.env.SetUserId(userId)
	return h.next.ServeHTTP(w, r)
}

func WithAuth(env *app.Env, next app.Handler) app.Handler {
	return authHandler{env: env, next: next}
}

func parseTokenUserId(token string, secret string) (string, error) {
	tokenService := newTokenService(secret)
	if err := tokenService.ParseToken(token); err != nil {
		return "", err
	}
	return tokenService.GetClaims().UserId, nil
}

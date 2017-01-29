package users

import (
	"dude_expenses/app"
	"dude_expenses/auth"
	"dude_expenses/passwords"
	"strconv"
)

type AuthService struct {
	env *app.Env
}

func NewAuthService(env *app.Env) *AuthService {
	return &AuthService{env: env}
}

func (service *AuthService) Authenticate(params UserParams) (*auth.Token, error) {
	var token *auth.Token

	user, err := service.findUser(params.Email)
	if err != nil {
		return token, err
	}

	if err = passwords.CheckPassword(user.EncryptedPassword, params.Password); err != nil {
		return token, &AuthInvalidPasswordError{}
	}

	token, err = auth.GenerateToken(strconv.FormatInt(user.Id, 10), service.env.GetAuthSecret())
	if err != nil {
		return token, err
	}
	return token, nil
}

func (service *AuthService) findUser(email string) (*User, error) {
	repository := NewRepository(service.env.GetDB())
	user, err := repository.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return user, &AuthInvalidEmailError{}
	}
	return user, nil
}

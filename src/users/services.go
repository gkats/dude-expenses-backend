package users

import (
	"app"
	"auth"
	"golang.org/x/crypto/bcrypt"
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

	if err = checkPassword(user.EncryptedPassword, params.Password); err != nil {
		return token, &AuthInvalidPasswordError{}
	}

	token, err = auth.GenerateToken(strconv.FormatInt(user.Id, 10))
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

func checkPassword(encryptedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}

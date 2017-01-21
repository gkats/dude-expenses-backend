package users

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int64  `json:"id,omitempty"`
	Email             string `json:"email,omitempty"`
	CreatedAt         string `json:"created_at,omitempty"`
	UpdatedAt         string `json:"updated_at,omitempty"`
	EncryptedPassword string `json:"-"`
}

func NewUser(params UserParams) (User, error) {
	user := User{Email: params.Email}
	err := user.encryptPassword(params.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (user *User) encryptPassword(password string) error {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		user.EncryptedPassword = string(encryptedPassword)
	}
	return err
}

type UserParams struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

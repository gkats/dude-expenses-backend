package users

import (
	"dude_expenses/passwords"
)

type User struct {
	Id                int64  `json:"id,omitempty"`
	Email             string `json:"email,omitempty"`
	CreatedAt         string `json:"createdAt,omitempty"`
	UpdatedAt         string `json:"updatedAt,omitempty"`
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
	encryptedPassword, err := passwords.GenerateEncrypted(password)
	if err == nil {
		user.EncryptedPassword = string(encryptedPassword)
	}
	return err
}

type UserParams struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

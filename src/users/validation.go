package users

import (
	"regexp"
)

type UserValidation struct {
	UserParams
	Errors     map[string][]string `json:"errors"`
	repository *Repository
}

func NewUserValidation(params UserParams, repository *Repository) UserValidation {
	return UserValidation{UserParams: params, repository: repository}
}

func (validation *UserValidation) IsValid() bool {
	validation.Errors = make(map[string][]string)

	if len(validation.Email) == 0 || !validation.isValidEmail() {
		validation.Errors["email"] = append(validation.Errors["email"], "must be a valid email address")
	}
	if validation.emailExists() {
		validation.Errors["email"] = append(validation.Errors["email"], "is already taken")
	}
	if len(validation.Password) < 6 {
		validation.Errors["password"] = append(validation.Errors["password"], "must be at least 6 characters")
	}

	return len(validation.Errors) == 0
}

func (validation UserValidation) isValidEmail() bool {
	matches, err := regexp.MatchString("^[^@]+@[^@]+\\.[^@]+$", validation.Email)
	return err == nil && matches
}

func (validation UserValidation) emailExists() bool {
	user, err := validation.repository.FindUserByEmail(validation.Email)
	return err != nil || user != nil
}

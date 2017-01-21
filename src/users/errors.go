package users

type AuthInvalidEmailError struct {
}

func (e *AuthInvalidEmailError) Error() string {
	return "Email does not exist"
}

type AuthInvalidPasswordError struct {
}

func (e *AuthInvalidPasswordError) Error() string {
	return "Password is invalid"
}

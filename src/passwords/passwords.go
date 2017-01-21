package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateEncrypted(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func CheckPassword(encryptedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
}
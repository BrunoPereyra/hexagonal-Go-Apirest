package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(passwordHash), errHash
}

func DecodePassword(PasswordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(PasswordHash), []byte(password))
	return err
}

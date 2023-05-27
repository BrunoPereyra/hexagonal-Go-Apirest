package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(passwordHash), errHash
}

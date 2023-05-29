package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string, passwordhashChan chan string) {
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(password), 10)
	if errHash != nil {
		passwordhashChan <- "errHash"
	} else {
		passwordhashChan <- string(passwordHash)
	}
}

func DecodePassword(PasswordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(PasswordHash), []byte(password))
	return err
}

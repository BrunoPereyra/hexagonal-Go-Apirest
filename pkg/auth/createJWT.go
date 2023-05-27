package auth

import (
	"hexagonal/internal/user/domain"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func CreateToken(user *domain.User) (string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("godotenv.Load JWT error")
	}
	TOKENPASSWORD := os.Getenv("TOKENPASSWORD")
	if TOKENPASSWORD == "" {
		TOKENPASSWORD = "udnwd38xm"
	}
	Claims := jwt.MapClaims{
		"_id":      user.ID,
		"username": user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	signedToken, errTokenSigned := token.SignedString([]byte(TOKENPASSWORD))
	return signedToken, errTokenSigned
}

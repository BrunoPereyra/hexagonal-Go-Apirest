package auth

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func parseToken(tokenString string) (*jwt.Token, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("godotenv error parseToken")
	}
	TOKENPASSWORD := os.Getenv("TOKENPASSWORD")
	if TOKENPASSWORD == "" {
		TOKENPASSWORD = "udnwd38xm"
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(TOKENPASSWORD), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	return token, nil
}

func ExtractDataFromToken(tokenString string) (string, string, error) {
	// Primero, parsea el token

	token, err := parseToken(tokenString)
	if err != nil {
		return "", "", err
	}
	// Luego, accede a los claims del token para obtener los datos que necesitas
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", fmt.Errorf("Invalid claims")
	}
	nameUser, ok := claims["nameuser"].(string)
	if !ok {
		return "", "", fmt.Errorf("Invalid nameUser")
	}
	_id, ok := claims["_id"].(string)
	if !ok {
		return "", "", fmt.Errorf("Invalid _id")
	}
	return nameUser, _id, nil
}

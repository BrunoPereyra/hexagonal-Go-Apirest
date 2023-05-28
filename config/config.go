package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func URI() string {
	if err := godotenv.Load(); err != nil {
		log.Println("godotenv.Load error URI mongo")
	}
	return os.Getenv("MONGODB_URI")

}
func CLOUDINARY_URL() string {
	if err := godotenv.Load(); err != nil {
		log.Println("godotenv.Load error CLOUDINARY_URL")
	}
	return os.Getenv("CLOUDINARY_URL")
}

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func URI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatal("godotenv.Load error URI mongo")
	}
	return os.Getenv("MONGODB_URI")

}

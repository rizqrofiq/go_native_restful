package helper

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Env(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

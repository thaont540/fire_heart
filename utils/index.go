package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Env(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return ""
	}

	return os.Getenv(key)
}

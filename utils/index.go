package utils

import (
	"github.com/joho/godotenv"
	"os"
)

func Env(key string) string {
	err := godotenv.Load()
	if err != nil {
		// for heroku
		return os.Getenv(key)
		//log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}

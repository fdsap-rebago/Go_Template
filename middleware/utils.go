// Package utils provides ...
package middleware

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func GetSQLEnv(key string) string {
	err := godotenv.Load(".env-SCRIPT")

	if err != nil {
		log.Fatalf("Error loading .env-SCRIPT file")
	}

	return os.Getenv(key)
}

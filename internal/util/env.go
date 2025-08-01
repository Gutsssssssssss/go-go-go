package util

import (
	"log"
	"os"
)

func EnsureEnvExist(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Error loading .env file: %s is not set", key)
	}
	return val
}

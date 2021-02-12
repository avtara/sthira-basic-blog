package helper

import (
	"os"

	"github.com/joho/godotenv"
)

//GetEnv read environtment variable if doesnt exist return fallback
func GetEnv(variable string, fallback string) string {
	godotenv.Load()
	response := os.Getenv(variable)
	if response != "" {
		return response
	}
	return fallback
}

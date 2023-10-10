package config

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}
}

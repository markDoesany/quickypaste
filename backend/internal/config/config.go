package config

import (
	"github.com/joho/godotenv"
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

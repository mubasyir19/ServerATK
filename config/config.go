package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Config() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Failed load .env file")
	}
}

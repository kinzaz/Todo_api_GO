package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Port: os.Getenv("SERVER_PORT"),
	}
}

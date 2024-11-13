package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	DSN string
}

type Config struct {
	Port string
	DB   DB
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file, using default config")
	}

	return &Config{
		Port: os.Getenv("SERVER_PORT"),
		DB: DB{
			DSN: os.Getenv("DSN"),
		},
	}
}

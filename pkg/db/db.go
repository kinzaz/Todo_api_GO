package db

import (
	"log"
	"todo_GO/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct{ *gorm.DB }

func NewDb() *DB {
	conf := configs.LoadConfig()

	db, err := gorm.Open(postgres.Open(conf.DB.DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed connection to DB: %v", err)
	}
	return &DB{db}
}

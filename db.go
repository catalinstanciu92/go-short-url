package main

import (
	"os"

	"github.com/go-pg/pg/v11"
	"github.com/joho/godotenv"
)

func connectToDB() *pg.DB {
	// load .env file
	godotenv.Load()
	// connect to the database
	db := pg.Connect(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	})

	return db
}

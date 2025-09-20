package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v11"
)

func generateShortURL(originalURL string) string {
	var db = connectToDB()
	ctx := context.Background()
	var id *int
	_, err := db.Query(ctx, pg.Scan(&id), "INSERT INTO urls (origin) VALUES (?) RETURNING id", originalURL)
	if err != nil {
		log.Fatalf("Could not insert into database: %s\n", err.Error())
	}

	return os.Getenv("BASE_URL") + "/url?q=" + fmt.Sprint(*id)
}

func getOriginalURL(id string) (string, error) {
	var db = connectToDB()
	ctx := context.Background()
	var originalURL *string

	_, err := db.QueryOne(ctx, pg.Scan(&originalURL), "SELECT origin FROM urls WHERE id = ?", id)
	if err != nil {
		return "", err
	}

	return *originalURL, nil
}

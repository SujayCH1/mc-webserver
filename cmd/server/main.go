package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"mc-webserver/internal/database"
)

func main() {

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	database.RunMigrations(dbURL)

	fmt.Println("Database ready")

	for {
		time.Sleep(10 * time.Second)
	}
}

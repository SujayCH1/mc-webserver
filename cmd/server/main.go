package main

import (
	"fmt"
	"log"

	"mc-webserver/internal/database"
	"mc-webserver/internal/router"
)

func main() {

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database ready")

	// Setup router
	r := router.SetUpRouter(db)

	// Start server
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

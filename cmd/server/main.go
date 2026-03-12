package main

import (
	"fmt"
	"log"

	"mc-webserver/internal/database"
	"mc-webserver/internal/router"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Database ready")

	r := router.SetUpRouter()

	r.Run(":8080")
}

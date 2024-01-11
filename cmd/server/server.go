package main

import (
	"log"

	"bizsearch/internal/database"
	"bizsearch/internal/http"
)

func main() {
	_, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	router := http.Create()
	router.Run("localhost:8080")
}

package main

import (
	"bizsearch/internal/http"
)

func main() {
	router := http.Create()
	router.Run("localhost:8080")
}

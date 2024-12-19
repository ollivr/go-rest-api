// main.go
package main

import (
	"log"
	"net/http"
	"restful-api/internal/handlers"
)

func main() {
	router := handlers.HandleRequests()
	log.Fatal(http.ListenAndServe(":10000", router))
}

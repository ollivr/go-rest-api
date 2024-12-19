// main.go
package main

import (
	"log"
	"net/http"
	"restful-api/internal/routes"
)

func main() {
	router := routes.HandleRequests()

	log.Fatal(http.ListenAndServe(":10000", router))

}

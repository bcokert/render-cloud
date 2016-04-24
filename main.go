package main

import (
	"log"
	"net/http"

	"github.com/bcokert/render-cloud/router"
)

// Initializes injection dependencies and then starts the web server
func main() {
	log.Println("Starting Server")

	// Create and run the web server
	r := router.CreateDefaultRouter()
	log.Fatal(http.ListenAndServe("localhost:8080/", r))
}

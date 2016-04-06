package main

import (
	"log"
	"net/http"

	"github.com/bcokert/render-cloud/router"
)

func main() {
	log.Println("Starting Server")

	r := router.CreateDefaultRouter()
	log.Fatal(http.ListenAndServe("localhost:8080/", r))
}

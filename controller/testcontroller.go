package controller

import (
	"net/http"
	"fmt"
	"html"
	"github.com/gorilla/mux"
	"log"
	"strconv"
)

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello, %q", html.EscapeString(request.URL.Path))
}

func TestEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Returning all test objects")
}

func TestXEndpoint(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	log.Printf("Request: %q", request.URL)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("Warning - Illegal id passed to /test/{id}")
		response.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(response, "That id cannot be found: %q", vars["id"])
	} else {
		fmt.Fprintf(response, "Returning test object %d", id)
	}
}

package controller

import (
	"encoding/json"
	"fmt"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/vector"
	"github.com/gorilla/mux"
	"github.com/lucasb-eyer/go-colorful"
	"html"
	"log"
	"net/http"
	"strconv"
)

//func PostRender(response http.ResponseWriter, request *http.Request) {
//	vars := mux.Vars(request)
//	log.Printf("Request: %q", request.URL)
//}

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

func SphereEndpoint(response http.ResponseWriter, request *http.Request) {
	log.Printf("Request: %q", request.URL)

	spheres := model.Spheres{
		model.Sphere{Origin: vector.Vector3{0, 0, 0}, Radius: 5.0, Color: colorful.Color{R: 1, G: 0, B: 0}},
		model.Sphere{Origin: vector.Vector3{600, 100, 0}, Radius: 7.0, Color: colorful.Color{R: 0, G: 1, B: 0}},
		model.Sphere{Origin: vector.Vector3{300, 400, 0}, Radius: 1.5, Color: colorful.Color{R: 0, G: 0, B: 1}},
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(spheres)
}

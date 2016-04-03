package router

import (
	"github.com/bcokert/render-cloud/controller"
	"github.com/gorilla/mux"
)

type Router interface {
	CreateRouter() *mux.Router
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controller.Index)
	router.HandleFunc("/test", controller.TestEndpoint)
	router.HandleFunc("/test/{id}", controller.TestXEndpoint)
	return router
}

package router

import (
	"github.com/gorilla/mux"
	"net/http"
)

type RouteConfig struct {
	Method           string
	Path             string
	ControllerAction http.HandlerFunc
}

func CreateDefaultRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	AddRoutesToRouter(router, RenderControllerRoutes)
	return router
}

func AddRoutesToRouter(router *mux.Router, routeConfigs []RouteConfig) {
	for _, routeConfig := range routeConfigs {
		router.
			Methods(routeConfig.Method).
			Path(routeConfig.Path).
			Name(routeConfig.Path).
			Handler(routeConfig.ControllerAction)
	}
}

package router

import "github.com/bcokert/render-cloud/controller"

var RenderControllerRoutes = []RouteConfig{
	RouteConfig{"GET", "/", controller.Index},
	RouteConfig{"GET", "/test", controller.TestEndpoint},
	RouteConfig{"GET", "/test/{id}", controller.TestXEndpoint},
	RouteConfig{"GET", "/spheres", controller.SphereEndpoint},
}

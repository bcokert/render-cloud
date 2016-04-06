package router

import "github.com/bcokert/render-cloud/controller"

var RenderControllerRoutes = []RouteConfig{
	RouteConfig{"POST", "/render", controller.PostRender},
}

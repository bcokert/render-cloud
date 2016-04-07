package router_test

import (
	"fmt"
	"github.com/bcokert/render-cloud/router"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func MakeHandler(varsChannel chan<- map[string]string, responseBody string) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, responseBody)
		varsChannel <- mux.Vars(request)
	}
}

func TestCreateDefaultRouter(t *testing.T) {
	expectedRoutes := []string{
		"/render",
	}
	expectedNumRoutes := len(expectedRoutes)

	muxRouter := router.CreateDefaultRouter()

	foundRoutes := map[string]bool{}
	muxRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		foundRoutes[route.GetName()] = true
		return nil
	})

	if len(foundRoutes) != expectedNumRoutes {
		t.Errorf("CreateDefaultRouter added %d routes, expected %d routes", len(foundRoutes), expectedNumRoutes)
	}

	for _, expectedRoute := range expectedRoutes {
		if _, ok := foundRoutes[expectedRoute]; ok != true {
			t.Errorf("CreateDefaultRouter did not have expected route %s", expectedRoute)
		}
	}
}

func TestAddRoutesToRouterEmptyRouteConfigsDoesNothing(t *testing.T) {
	routeConfigs := []router.RouteConfig{}

	muxRouter := mux.NewRouter()
	router.AddRoutesToRouter(muxRouter, routeConfigs)

	routeCount := 0
	muxRouter.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		routeCount++
		return nil
	})

	if routeCount > 0 {
		t.Errorf("AddRoutesToRouter added %d routes when given an empty route config", routeCount)
	}
}

func TestAddRoutesToRouterAddsRouteToRouterCorrectly(t *testing.T) {
	varsChannel := make(chan map[string]string, 3)
	routeConfigs := []router.RouteConfig{
		router.RouteConfig{"GET", "/test1", MakeHandler(varsChannel, "response1")},
		router.RouteConfig{"POST", "/test2/path234/", MakeHandler(varsChannel, "response2")},
		router.RouteConfig{"DELETE", "/test3/{id}", MakeHandler(varsChannel, "response3")},
	}

	muxRouter := mux.NewRouter()
	router.AddRoutesToRouter(muxRouter, routeConfigs)

	testCases := []struct {
		Url         string
		Vars        map[string]string
		Body        string
		RouteConfig router.RouteConfig
	}{
		{"/test1", map[string]string{}, "response1", routeConfigs[0]},
		{"/test2/path234/", map[string]string{}, "response2", routeConfigs[1]},
		{"/test3/52", map[string]string{"id": "52"}, "response3", routeConfigs[2]},
	}

	for _, testCase := range testCases {
		testutils.ExpectRouterRoutes(t, muxRouter, testCase.RouteConfig.Method, testCase.Url, nil, 200, testCase.Body, varsChannel, testCase.Vars)
	}
}

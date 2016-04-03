package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/bcokert/render-cloud/router"
)

func TestGetTestXSucceed(t *testing.T) {
	request, err := http.NewRequest("GET", "/test/3311", nil)
	if err != nil {
		t.Errorf("Failed to create Mock Request")
	}
	response := httptest.NewRecorder()

	router.CreateRouter().ServeHTTP(response, request)

	if (response.Code != 200) {
		t.Errorf("GET test/3311 had status %d, expected %d", response.Code, 200)
	}

	if (response.Body.String() != "Returning test object 3311") {
		t.Errorf("GET test/3311 had body %q, expected %q", response.Body.String(), "Returning test object 3311")
	}
}

func TestGetTestXFail(t *testing.T) {
	request, err := http.NewRequest("GET", "/test/4df342a", nil)
	if err != nil {
		t.Errorf("Failed to create Mock Request")
	}
	response := httptest.NewRecorder()

	router.CreateRouter().ServeHTTP(response, request)

	if (response.Code != 400) {
		t.Errorf("GET test/4df342a had status %d, expected %d", response.Code, 404)
	}

	if (response.Body.String() != "That id cannot be found: \"4df342a\"") {
		t.Errorf("GET test/4df342a had body %q, expected %q", response.Body.String(), "That id cannot be found: \"4df342a\"")
	}
}

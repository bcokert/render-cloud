package testutils

import (
	"github.com/bcokert/render-cloud/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ExpectControllerRequestString(t *testing.T, method, url string, expectedStatusCode int, expectedBody string) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		t.Errorf("Failed to create Mock Request for %s %s test", method, url)
	}

	response := httptest.NewRecorder()

	router.CreateRouter().ServeHTTP(response, request)

	if response.Code != expectedStatusCode {
		t.Errorf("%s %s had status %d, expected %d", method, url, response.Code, expectedStatusCode)
	}

	if response.Body.String() != expectedBody {
		t.Errorf("%s %s had body %q, expected %q", method, url, response.Body.String(), expectedBody)
	}
}

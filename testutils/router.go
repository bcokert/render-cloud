package testutils

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"
	"github.com/bcokert/render-cloud/model"
)

func ExpectRouterRoutes(t *testing.T, router *mux.Router, method, url string, body interface{}, expectedStatusCode int, expectedBody string, varsChannel <-chan map[string]string, expectedVars map[string]string) {
	var request *http.Request
	var err error

	if body != nil {
		jsonString, err := model.ToJson(model.DefaultMarshaler, body)
		if err != nil {
			t.Errorf("Failed to encode provided body into json string")
		}

		request, err = http.NewRequest(method, url, bytes.NewReader([]byte(jsonString)))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		t.Errorf("Failed to create Mock Request for %s %s test", method, url)
	}

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code != expectedStatusCode {
		t.Errorf("%s %s had status %d, expected %d", method, url, response.Code, expectedStatusCode)
	}

	if response.Body.String() != expectedBody {
		t.Errorf("%s %s had body %q, expected %q", method, url, response.Body.String(), expectedBody)
	}

	if varsChannel != nil && expectedVars != nil {
		vars := <-varsChannel // swallow the vars even if there are none. Thus we expect an empty map to always be passed
		if len(expectedVars) > 0 {
			if len(vars) == 0 {
				t.Errorf("Handler for %s %s did not passed an empty channel into the vars channel", method, url)
			}
			for expectedVarName, expectedVarValue := range expectedVars {
				actualVarValue, ok := vars[expectedVarName]
				if ok != true {
					t.Errorf("%s %s is missing expected var %s", method, url, expectedVarName)
				}

				if expectedVarValue != actualVarValue {
					t.Errorf("%s %s var %s has value '%s', expcted '%s'", method, url, expectedVarName, actualVarValue, expectedVarValue)
				}
			}
		}
	}
}

package controller_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"net/http"
	"testing"
)

func TestGetTestXSucceed(t *testing.T) {
	testutils.ExpectControllerRequestString(t, http.MethodGet, "/test/3311", http.StatusOK, "Returning test object 3311")
}

func TestGetTestXFail(t *testing.T) {
	testutils.ExpectControllerRequestString(t, http.MethodGet, "/test/4df342a", http.StatusBadRequest, "That id cannot be found: \"4df342a\"")
}

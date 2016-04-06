package controller_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"net/http"
	"testing"
	"github.com/gorilla/mux"
	"github.com/bcokert/render-cloud/router"
)

func createRouterForTest() *mux.Router {
	return router.CreateDefaultRouter()
}

func TestPostRenderSucceed(t *testing.T) {
	testutils.ExpectRouterRoutes(t, createRouterForTest(), http.MethodPost, "/render", http.StatusOK, "", nil, nil)
}

package controller_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"net/http"
	"testing"
	"github.com/bcokert/render-cloud/router"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/utils"
)

func TestPostRenderSucceed(t *testing.T) {
	r := router.CreateDefaultRouter()
	testutils.ExpectRouterRoutes(t, r, http.MethodPost, "/render", model.Scene{utils.UintPointer(2345), nil, nil}, http.StatusOK, "", nil, nil)
}

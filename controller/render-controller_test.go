package controller_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/router"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/utils"
	"net/http"
	"testing"
)

func TestPostRenderSucceed(t *testing.T) {
	r := router.CreateDefaultRouter()
	scene := model.Scene{utils.UintPointer(2345), nil, nil}
	sceneJson, err := model.ToJson(model.DefaultMarshaler, scene)
	if err != nil {
		t.Errorf("Failed to convert scene to json, to verify response from server: %s", err.Error())
	}
	testutils.ExpectRouterRoutes(t, r, http.MethodPost, "/render", scene, http.StatusOK, sceneJson+"\n", nil, nil)
}

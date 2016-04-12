package controller_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/router"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/utils"
	"net/http"
	"testing"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/bcokert/render-cloud/model/materials"
)

func TestPostRenderSucceed(t *testing.T) {
	r := router.CreateDefaultRouter()
	scene := model.Scene{
		utils.UintPointer(2345),
		&model.World{
			Ambient: &colorful.Color{0.2,0.2,0.2},
			Background: &colorful.Color{0,0,0},
			Camera: &model.Camera{
				Origin: &mgl64.Vec3{0,0,-8},
				Direction: &mgl64.Vec3{0,0,1},
				Up: &mgl64.Vec3{0,1,0},
				ScreenWidth: utils.FloatPointer(5),
				ScreenHeight: utils.FloatPointer(5),
				ScreenDistance: utils.FloatPointer(2),
			},
			Lights: &[]model.Light{
				model.Light{
					Direction: &mgl64.Vec3{1, -1, 1},
					Color: &colorful.Color{0.4, 0.4, 0.4},
				},
			},
		},
		&[]primitives.Sphere{
			primitives.Sphere{
				Origin: &mgl64.Vec3{0,0,0},
				Radius: utils.FloatPointer(4),
				Material: &materials.Material{
					Color: &colorful.Color{0.7,0.1,0.1},
					Shininess: utils.FloatPointer(10),
				},
			},
		}}
	sceneJson, err := model.ToJson(model.DefaultMarshaler, scene)
	if err != nil {
		t.Errorf("Failed to convert scene to json, to verify response from server: %s", err.Error())
	}

	testutils.ExpectRouterRoutes(t, r, http.MethodPost, "/render", scene, http.StatusOK, sceneJson+"\n", nil, nil)
}

package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/bcokert/render-cloud/utils"
)

func TestCameraJsonEncodes(t *testing.T) {
	camera := model.Camera{
		Origin: &mgl64.Vec3{10, 10, -50},
		Direction: &mgl64.Vec3{0, 0, 1},
		Up: &mgl64.Vec3{0, 1, 0},
		ScreenWidth: utils.FloatPointer(4),
		ScreenHeight: utils.FloatPointer(3),
		ScreenDistance: utils.FloatPointer(1),
	}

	expectedJson := "{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0],\"screenWidth\":4,\"screenHeight\":3,\"screenDistance\":1}"

	testutils.ExpectJsonEncoding(t, &camera, expectedJson)
}

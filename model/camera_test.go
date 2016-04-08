package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
)

func TestCameraJsonEncodes(t *testing.T) {
	camera := model.Camera{Origin: &mgl64.Vec3{10, 10, -50}, Direction: &mgl64.Vec3{0, 0, 1}, Up: &mgl64.Vec3{0, 1, 0}}

	expectedJson := "{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0]}"

	testutils.ExpectJsonEncoding(t, &camera, expectedJson)
}

package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/vector"
	"testing"
)

func TestCameraJsonEncodes(t *testing.T) {
	camera := model.Camera{Origin: vector.Vector3{10, 10, -50}, Direction: vector.Vector3{0, 0, 1}, Up: vector.Vector3{0, 1, 0}}

	expectedJson := "{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0]}"

	testutils.ExpectJsonEncoding(t, &camera, expectedJson)
}

package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
)

func TestWorldJsonEncodes(t *testing.T) {
	world := model.World{
		Ambient:    &colorful.Color{1, 0, 0},
		Background: &colorful.Color{1, 1, 1},
		Camera:     &model.Camera{Origin: &mgl64.Vec3{10, 10, -50}, Direction: &mgl64.Vec3{0, 0, 1}, Up: &mgl64.Vec3{0, 1, 0}},
	}

	expectedJson := "{\"ambient\":{\"R\":1,\"G\":0,\"B\":0},\"background\":{\"R\":1,\"G\":1,\"B\":1},\"camera\":{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0]}}"

	testutils.ExpectJsonEncoding(t, &world, expectedJson)
}

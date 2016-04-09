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
		Lights:     &[]model.Light{
			{Direction: &mgl64.Vec3{-0.5,-0.5,-1}, Color: &colorful.Color{0.5,0.5,0.4}},
			{Direction: &mgl64.Vec3{0.5,0.5,0}, Color: &colorful.Color{0.1,0.1,0.2}},
		},
	}

	expectedJson := "{\"ambient\":{\"R\":1,\"G\":0,\"B\":0},\"background\":{\"R\":1,\"G\":1,\"B\":1},\"camera\":{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0]},\"lights\":[{\"direction\":[-0.5,-0.5,-1],\"color\":{\"R\":0.5,\"G\":0.5,\"B\":0.4}},{\"direction\":[0.5,0.5,0],\"color\":{\"R\":0.1,\"G\":0.1,\"B\":0.2}}]}"

	testutils.ExpectJsonEncoding(t, &world, expectedJson)
}

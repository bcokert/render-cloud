package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
)

func TestLightJsonEncodes(t *testing.T) {
	light := model.Light{Direction: &mgl64.Vec3{0, 0, 1}, Color: &colorful.Color{0.5,0.5,0.4}}

	expectedJson := "{\"direction\":[0,0,1],\"color\":{\"R\":0.5,\"G\":0.5,\"B\":0.4}}"

	testutils.ExpectJsonEncoding(t, &light, expectedJson)
}

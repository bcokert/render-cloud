package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/utils"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/bcokert/render-cloud/model/materials"
)

func TestSceneJsonEncodes(t *testing.T) {
	scene := model.Scene{
		Id: 8752,
		World: model.World{
			Ambient:    &colorful.Color{1, 0, 0},
			Background: &colorful.Color{1, 1, 1},
			Camera:     &model.Camera{Origin: &mgl64.Vec3{10, 10, -50}, Direction: &mgl64.Vec3{0, 0, 1}, Up: &mgl64.Vec3{0, 1, 0}},
		},
		Spheres: []primitives.Sphere{
			primitives.Sphere{Origin: &mgl64.Vec3{0, 0, 0}, Radius: utils.FloatPointer(5.0), Material: &materials.Material{Color: &colorful.Color{R: 1, G: 0, B: 0}, Shininess: utils.FloatPointer(2)}},
			primitives.Sphere{Origin: &mgl64.Vec3{600, 100, 0}, Radius: utils.FloatPointer(7.0), Material: &materials.Material{Color: &colorful.Color{R: 0, G: 1, B: 0}, Shininess: utils.FloatPointer(1)}},
			primitives.Sphere{Origin: &mgl64.Vec3{300, 400, 0}, Radius: utils.FloatPointer(1.5), Material: &materials.Material{Color: &colorful.Color{R: 0, G: 0, B: 1}, Shininess: utils.FloatPointer(6)}},
		},
	}

	expectedJson := "{\"id\":8752,\"world\":{\"ambient\":{\"R\":1,\"G\":0,\"B\":0},\"background\":{\"R\":1,\"G\":1,\"B\":1},\"camera\":{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0]}},\"spheres\":[{\"origin\":[0,0,0],\"radius\":5,\"material\":{\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":2}},{\"origin\":[600,100,0],\"radius\":7,\"material\":{\"color\":{\"R\":0,\"G\":1,\"B\":0},\"shininess\":1}},{\"origin\":[300,400,0],\"radius\":1.5,\"material\":{\"color\":{\"R\":0,\"G\":0,\"B\":1},\"shininess\":6}}]}"
	testutils.ExpectJsonEncoding(t, &scene, expectedJson)
}

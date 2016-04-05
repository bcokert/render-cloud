package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/vector"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
)

func TestSceneJsonEncodes(t *testing.T) {
	scene := model.Scene{
		Id: 8752,
		World: model.World{
			Ambient:    colorful.Color{1, 0, 0},
			Background: colorful.Color{1, 1, 1},
			Camera:     model.Camera{Origin: vector.Vector3{10, 10, -50}, Direction: vector.Vector3{0, 0, 1}, Up: vector.Vector3{0, 1, 0}},
		},
		Spheres: model.Spheres{
			model.Sphere{Origin: vector.Vector3{0, 0, 0}, Radius: 5.0, Color: colorful.Color{R: 1, G: 0, B: 0}, Shininess: 2},
			model.Sphere{Origin: vector.Vector3{600, 100, 0}, Radius: 7.0, Color: colorful.Color{R: 0, G: 1, B: 0}, Shininess: 1},
			model.Sphere{Origin: vector.Vector3{300, 400, 0}, Radius: 1.5, Color: colorful.Color{R: 0, G: 0, B: 1}, Shininess: 6},
		},
	}

	expectedJson := "{" +
		"\"id\":8752," +
		"\"world\":{\"ambient\":{\"R\":1,\"G\":0,\"B\":0},\"background\":{\"R\":1,\"G\":1,\"B\":1},\"camera\":{\"origin\":[10,10,-50],\"direction\":[0,0,1],\"up\":[0,1,0]}}," +
		"\"spheres\":[{\"origin\":[0,0,0],\"radius\":5,\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":2},{\"origin\":[600,100,0],\"radius\":7,\"color\":{\"R\":0,\"G\":1,\"B\":0},\"shininess\":1},{\"origin\":[300,400,0],\"radius\":1.5,\"color\":{\"R\":0,\"G\":0,\"B\":1},\"shininess\":6}]" +
		"}"

	testutils.ExpectJsonEncoding(t, &scene, expectedJson)
}

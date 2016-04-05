package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/vector"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
)

func TestSphereJsonEncodes(t *testing.T) {
	sphere := model.Sphere{Origin: vector.Vector3{0, 0, 0}, Radius: 5.0, Color: colorful.Color{R: 1, G: 0, B: 0}}

	expectedJson := "{\"origin\":[0,0,0],\"radius\":5,\"color\":{\"R\":1,\"G\":0,\"B\":0}}"

	testutils.ExpectJsonEncoding(t, &sphere, expectedJson)
}

func TestSpheresJsonEncodes(t *testing.T) {
	spheres := model.Spheres{
		model.Sphere{Origin: vector.Vector3{0, 0, 0}, Radius: 5.0, Color: colorful.Color{R: 1, G: 0, B: 0}},
		model.Sphere{Origin: vector.Vector3{600, 100, 0}, Radius: 7.0, Color: colorful.Color{R: 0, G: 1, B: 0}},
		model.Sphere{Origin: vector.Vector3{300, 400, 0}, Radius: 1.5, Color: colorful.Color{R: 0, G: 0, B: 1}},
	}

	expectedJson := "[{\"origin\":[0,0,0],\"radius\":5,\"color\":{\"R\":1,\"G\":0,\"B\":0}},{\"origin\":[600,100,0],\"radius\":7,\"color\":{\"R\":0,\"G\":1,\"B\":0}},{\"origin\":[300,400,0],\"radius\":1.5,\"color\":{\"R\":0,\"G\":0,\"B\":1}}]"

	testutils.ExpectJsonEncoding(t, &spheres, expectedJson)
}

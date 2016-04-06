package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/vector"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/bcokert/render-cloud/utils"
)

func TestSphereJsonEncodes(t *testing.T) {
	sphere := model.Sphere{Origin: &vector.Vector3{0, 0, 0}, Radius: utils.FloatPointer(5.0), Color: &colorful.Color{R: 1, G: 0, B: 0}, Shininess: utils.FloatPointer(9)}

	expectedJson := "{\"origin\":[0,0,0],\"radius\":5,\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":9}"

	testutils.ExpectJsonEncoding(t, &sphere, expectedJson)
}

func TestSpheresJsonEncodes(t *testing.T) {
	spheres := model.Spheres{
		model.Sphere{Origin: &vector.Vector3{0, 0, 0}, Radius: utils.FloatPointer(5.0), Color: &colorful.Color{R: 1, G: 0, B: 0}, Shininess: utils.FloatPointer(1)},
		model.Sphere{Origin: &vector.Vector3{600, 100, 0}, Radius: utils.FloatPointer(7.0), Color: &colorful.Color{R: 0, G: 1, B: 0}, Shininess: utils.FloatPointer(4)},
		model.Sphere{Origin: &vector.Vector3{300, 400, 0}, Radius: utils.FloatPointer(1.5), Color: &colorful.Color{R: 0, G: 0, B: 1}, Shininess: utils.FloatPointer(1.23)},
	}

	expectedJson := "[{\"origin\":[0,0,0],\"radius\":5,\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":1},{\"origin\":[600,100,0],\"radius\":7,\"color\":{\"R\":0,\"G\":1,\"B\":0},\"shininess\":4},{\"origin\":[300,400,0],\"radius\":1.5,\"color\":{\"R\":0,\"G\":0,\"B\":1},\"shininess\":1.23}]"

	testutils.ExpectJsonEncoding(t, &spheres, expectedJson)
}

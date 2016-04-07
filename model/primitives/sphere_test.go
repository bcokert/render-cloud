package primitives_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/utils"
	"github.com/bcokert/render-cloud/vector"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/bcokert/render-cloud/model/materials"
)

func TestSphereJsonEncodes(t *testing.T) {
	sphere := primitives.Sphere{Origin: &vector.Vector3{0, 0, 0}, Radius: utils.FloatPointer(5.0), Material: &materials.Material{Color: &colorful.Color{R: 1, G: 0, B: 0}, Shininess: utils.FloatPointer(9)}}

	expectedJson := "{\"origin\":[0,0,0],\"radius\":5,\"material\":{\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":9}}"

	testutils.ExpectJsonEncoding(t, &sphere, expectedJson)
}

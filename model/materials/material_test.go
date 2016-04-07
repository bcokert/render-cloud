package materials_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/utils"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/bcokert/render-cloud/model/materials"
)

func TestMaterialJsonEncodes(t *testing.T) {
	material := materials.Material{Color: &colorful.Color{R: 1, G: 0, B: 0}, Shininess: utils.FloatPointer(9)}

	expectedJson := "{\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":9}"

	testutils.ExpectJsonEncoding(t, &material, expectedJson)
}

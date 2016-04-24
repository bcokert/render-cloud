package materials_test

import (
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"encoding/json"
	"github.com/bcokert/render-cloud/validation"
)

func TestMaterialJsonEncodes(t *testing.T) {
	material := materials.Material{Color: colorful.Color{1, 0, 0}, Shininess: 9}

	expectedJson := `{"color":{"R":1,"G":0,"B":0},"shininess":9}`

	testutils.ExpectJsonEncoding(t, &material, expectedJson)
}

func TestMaterial_FromPostRequest(t *testing.T) {
	validator := validation.NewValidator()

	testCases := map[string]struct {
		PostRequestJson string
		Expected        materials.Material
	}{
		"empty material": {
			PostRequestJson: `{}`,
			Expected:        materials.Material{Color: colorful.Color{0.8, 0.1, 0.1}, Shininess: 1},
		},
		"partial material": {
			PostRequestJson: `{"color": {"R":0,"G":1,"B":0}}`,
			Expected:        materials.Material{Color: colorful.Color{0, 1, 0}, Shininess: 1},
		},
		"custom material": {
			PostRequestJson: `{"shininess":42, "color":{"R":0,"G":1,"B":0}}`,
			Expected:        materials.Material{Color: colorful.Color{0, 1, 0}, Shininess: 42},
		},
	}

	validationTestCases := map[string]struct {
		PostRequestJson string
		Expected string
	}{
		"Invalid color": {
			PostRequestJson: `{"color": {"R": 0, "B": 0, "G": 2}}`,
			Expected: "Color: not a valid color (rgb in 0-1)",
		},
		"Shininess 0": {
			PostRequestJson: `{"shininess": 0}`,
			Expected: "Shininess: zero value",
		},
		"Shininess <0": {
			PostRequestJson: `{"shininess": -3}`,
			Expected: "Shininess: less than min",
		},
	}

	for name, testCase := range testCases {
		var postRequest materials.MaterialPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var material materials.Material
		if err := material.FromPostRequest(validator, postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error was returned: %s", name, err.Error())
			continue
		}

		if material != testCase.Expected {
			t.Errorf("'%s' failed. Expected %#v, received %#v", name, testCase.Expected, material)
		}
	}

	for name, testCase := range validationTestCases {
		var postRequest materials.MaterialPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var material materials.Material
		if err := material.FromPostRequest(validator, postRequest); err == nil || err.Error() != testCase.Expected {
			t.Errorf("'%s' failed. Expected error %s, received %v", name, testCase.Expected, err)
		}
	}
}

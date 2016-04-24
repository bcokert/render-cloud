package model_test

import (
	"encoding/json"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/validation"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
)

func TestLight_JsonEncodes(t *testing.T) {
	light := model.Light{Direction: mgl64.Vec3{0, 0, 1}, Color: colorful.Color{0.5, 0.5, 0.4}}

	expectedJson := `{"direction":[0,0,1],"color":{"R":0.5,"G":0.5,"B":0.4}}`

	testutils.ExpectJsonEncoding(t, &light, expectedJson)
}

func TestLight_FromPostRequest(t *testing.T) {
	validator := validation.NewValidator()

	testCases := map[string]struct {
		PostRequestJson string
		Expected        model.Light
	}{
		"empty light": {
			PostRequestJson: `{}`,
			Expected:        model.Light{Direction: mgl64.Vec3{-1, -1, 1}, Color: colorful.Color{0.7, 0.7, 0.7}},
		},
		"partial light": {
			PostRequestJson: `{"color": {"R":0,"G":1,"B":0}}`,
			Expected:        model.Light{Direction: mgl64.Vec3{-1, -1, 1}, Color: colorful.Color{0, 1, 0}},
		},
		"custom light": {
			PostRequestJson: `{"direction":[-3,4,6], "color":{"R":0,"G":1,"B":0}}`,
			Expected:        model.Light{Direction: mgl64.Vec3{-3, 4, 6}, Color: colorful.Color{0, 1, 0}},
		},
	}

	validationTestCases := map[string]struct {
		PostRequestJson string
		Expected string
	}{
		"Invalid direction": {
			PostRequestJson: `{"direction": [0,0,0]}`,
			Expected: "Direction: not a direction vector (found [0,0,0])",
		},
		"Invalid color": {
			PostRequestJson: `{"color": {"R": 0, "B": 0, "G": 2}}`,
			Expected: "Color: not a valid color (rgb in 0-1)",
		},
	}

	for name, testCase := range testCases {
		var postRequest model.LightPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var light model.Light
		if err := light.FromPostRequest(validator, postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error was returned: %s", name, err.Error())
			continue
		}

		if light != testCase.Expected {
			t.Errorf("'%s' failed. Expected %#v, received %#v", name, testCase.Expected, light)
		}
	}

	for name, testCase := range validationTestCases {
		var postRequest model.LightPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var light model.Light
		if err := light.FromPostRequest(validator, postRequest); err == nil || err.Error() != testCase.Expected {
			t.Errorf("'%s' failed. Expected error %s, received %v", name, testCase.Expected, err)
		}
	}
}

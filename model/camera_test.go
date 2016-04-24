package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
	"encoding/json"
	"github.com/bcokert/render-cloud/validation"
)

func TestCameraJsonEncodes(t *testing.T) {
	camera := model.Camera{
		Origin: mgl64.Vec3{10, 10, -50},
		Direction: mgl64.Vec3{0, 0, 1},
		Up: mgl64.Vec3{0, 1, 0},
		ScreenWidth: 4,
		ScreenHeight: 3,
		ScreenDistance: 1,
	}

	expectedJson := `{"origin":[10,10,-50],"direction":[0,0,1],"up":[0,1,0],"screenWidth":4,"screenHeight":3,"screenDistance":1}`

	testutils.ExpectJsonEncoding(t, &camera, expectedJson)
}

func TestCamera_FromPostRequest(t *testing.T) {
	validator := validation.NewValidator()

	testCases := map[string]struct {
		PostRequestJson string
		Expected        model.Camera
	}{
		"empty camera": {
			PostRequestJson: `{}`,
			Expected:        model.Camera{Origin: mgl64.Vec3{0, 0, -10}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 4, ScreenHeight: 3, ScreenDistance: 1},
		},
		"partial camera": {
			PostRequestJson: `{"screenWidth":100, "screenHeight": 200}`,
			Expected:        model.Camera{Origin: mgl64.Vec3{0, 0, -10}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 100, ScreenHeight: 200, ScreenDistance: 1},
		},
		"custom camera": {
			PostRequestJson: `{"origin": [0,0,0], "up": [1,1,1], "direction": [1,2,3], "screenWidth": 55, "screenHeight": 44, "screenDistance": 22}`,
			Expected:        model.Camera{Origin: mgl64.Vec3{0, 0, 0}, Direction: mgl64.Vec3{1, 2, 3}, Up: mgl64.Vec3{1, 1, 1}, ScreenWidth: 55, ScreenHeight: 44, ScreenDistance: 22},
		},
	}

	validationTestCases := map[string]struct {
		PostRequestJson string
		Expected string
	}{
		"Invalid Direction": {
			PostRequestJson: `{"direction": [0,0,0]}`,
			Expected: "Direction: not a direction vector (found [0,0,0])",
		},
		"Invalid Up": {
			PostRequestJson: `{"up": [0,0,0]}`,
			Expected: "Up: not a direction vector (found [0,0,0])",
		},
		"ScreenWidth min=0": {
			PostRequestJson: `{"screenWidth": -4}`,
			Expected: "ScreenWidth: less than min",
		},
		"ScreenHeight min=0": {
			PostRequestJson: `{"screenHeight": -4}`,
			Expected: "ScreenHeight: less than min",
		},
		"ScreenDistance min=0": {
			PostRequestJson: `{"screenDistance": -4}`,
			Expected: "ScreenDistance: less than min",
		},
		"ScreenWidth nonzero": {
			PostRequestJson: `{"screenWidth": 0}`,
			Expected: "ScreenWidth: zero value",
		},
		"ScreenHeight nonzero": {
			PostRequestJson: `{"screenHeight": 0}`,
			Expected: "ScreenHeight: zero value",
		},
		"ScreenDistance nonzero": {
			PostRequestJson: `{"screenDistance": 0}`,
			Expected: "ScreenDistance: zero value",
		},
	}

	for name, testCase := range testCases {
		var postRequest model.CameraPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var camera model.Camera
		if err := camera.FromPostRequest(validator, postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error was returned: %s", name, err.Error())
			continue
		}

		if camera != testCase.Expected {
			t.Errorf("'%s' failed. Expected %#v, received %#v", name, testCase.Expected, camera)
		}
	}

	for name, testCase := range validationTestCases {
		var postRequest model.CameraPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var camera model.Camera
		if err := camera.FromPostRequest(validator, postRequest); err == nil || err.Error() != testCase.Expected {
			t.Errorf("'%s' failed. Expected error %s, received %v", name, testCase.Expected, err)
		}
	}
}

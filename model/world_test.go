package model_test

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/go-gl/mathgl/mgl64"
	"encoding/json"
	"github.com/bcokert/render-cloud/validation"
)

func TestWorldJsonEncodes(t *testing.T) {
	world := model.World{
		Ambient:    colorful.Color{1, 0, 0},
		Background: colorful.Color{1, 1, 1},
		Camera:     model.Camera{Origin: mgl64.Vec3{10, 10, -50}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 3, ScreenHeight: 10, ScreenDistance: 44},
		Lights:     []model.Light{
			model.Light{Direction: mgl64.Vec3{-0.5,-0.5,-1}, Color: colorful.Color{0.5,0.5,0.4}},
			model.Light{Direction: mgl64.Vec3{0.5,0.5,0}, Color: colorful.Color{0.1,0.1,0.2}},
		},
	}

	expectedJson := `{"ambient":{"R":1,"G":0,"B":0},"background":{"R":1,"G":1,"B":1},"camera":{"origin":[10,10,-50],"direction":[0,0,1],"up":[0,1,0],"screenWidth":3,"screenHeight":10,"screenDistance":44},"lights":[{"direction":[-0.5,-0.5,-1],"color":{"R":0.5,"G":0.5,"B":0.4}},{"direction":[0.5,0.5,0],"color":{"R":0.1,"G":0.1,"B":0.2}}]}`

	testutils.ExpectJsonEncoding(t, &world, expectedJson)
}

func TestWorld_FromPostRequest(t *testing.T) {
	validator := validation.NewValidator()

	testCases := map[string]struct {
		PostRequestJson string
		Expected        model.World
	}{
		"empty world": {
			PostRequestJson: `{}`,
			Expected:        model.World{
				Ambient:    colorful.Color{0.1, 0.1, 0.1},
				Background: colorful.Color{0, 0, 0},
				Camera:     model.Camera{Origin: mgl64.Vec3{0, 0, -10}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 4, ScreenHeight: 3, ScreenDistance: 1},
				Lights:     []model.Light{},
			},
		},
		"partial world": {
			PostRequestJson: `{"lights": [{"direction": [1,1,1]}]}`,
			Expected:        model.World{
				Ambient:    colorful.Color{0.1, 0.1, 0.1},
				Background: colorful.Color{0, 0, 0},
				Camera:     model.Camera{Origin: mgl64.Vec3{0, 0, -10}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 4, ScreenHeight: 3, ScreenDistance: 1},
				Lights:     []model.Light{
					model.Light{Direction: mgl64.Vec3{1,1,1}, Color: colorful.Color{0.7,0.7,0.7}},
				},
			},
		},
		"custom world": {
			PostRequestJson: `{"ambient": {"R":1, "G":0, "B":1}, "background": {"R": 1, "G": 1, "B": 1}, "camera": {"origin": [0,0,0], "up": [1,1,1], "direction": [1,2,3], "screenWidth": 55, "screenHeight": 44, "screenDistance": 22}, "lights": [{"direction": [1,1,1], "color": {"R":0, "G": 1, "B": 0}}]}`,
			Expected:        model.World{
				Ambient:    colorful.Color{1, 0, 1},
				Background: colorful.Color{1, 1, 1},
				Camera:     model.Camera{Origin: mgl64.Vec3{0, 0, 0}, Direction: mgl64.Vec3{1, 2, 3}, Up: mgl64.Vec3{1, 1, 1}, ScreenWidth: 55, ScreenHeight: 44, ScreenDistance: 22},
				Lights:     []model.Light{
					model.Light{Direction: mgl64.Vec3{1,1,1}, Color: colorful.Color{0,1,0}},
				},
			},
		},
	}

	validationTestCases := map[string]struct {
		PostRequestJson string
		Expected string
	}{
		"Invalid ambient": {
			PostRequestJson: `{"ambient": {"R": 2, "G": 0, "B": 2}}`,
			Expected: "Ambient: not a valid color (rgb in 0-1)",
		},
		"Invalid background": {
			PostRequestJson: `{"background": {"R": 0, "B": 3, "G": 0}}`,
			Expected: "Background: not a valid color (rgb in 0-1)",
		},
		"Invalid camera": {
			PostRequestJson: `{"camera": {"up": [0,0,0]}}`,
			Expected: "Camera: not a valid Camera (Up: not a direction vector (found [0,0,0]))",
		},
		"Invalid Lights": {
			PostRequestJson: `{"lights": [{"direction": [0,0,0]}]}`,
			Expected: "Lights: not a valid []Light (Direction: not a direction vector (found [0,0,0]))",
		},
	}

	for name, testCase := range testCases {
		var postRequest model.WorldPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var world model.World
		if err := world.FromPostRequest(validator, postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error was returned: %s", name, err.Error())
			continue
		}

		if !world.Equals(testCase.Expected) {
			t.Errorf("'%s' failed. Expected %#v, received %#v", name, testCase.Expected, world)
		}
	}

	for name, testCase := range validationTestCases {
		var postRequest model.WorldPostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var world model.World
		if err := world.FromPostRequest(validator, postRequest); err == nil || err.Error() != testCase.Expected {
			t.Errorf("'%s' failed. Expected error %s, received %v", name, testCase.Expected, err)
		}
	}
}

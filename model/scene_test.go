package model_test

import ()
import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/bcokert/render-cloud/testutils"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/bcokert/render-cloud/validation"
	"encoding/json"
)

func TestSceneJsonEncodes(t *testing.T) {
	scene := model.Scene{
		Id: 8752,
		World: model.World{
			Ambient:    colorful.Color{1, 0, 0},
			Background: colorful.Color{1, 1, 1},
			Camera:     model.Camera{Origin: mgl64.Vec3{10, 10, -50}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 2, ScreenHeight: 4, ScreenDistance: 11},
			Lights:     []model.Light{
				model.Light{Direction: mgl64.Vec3{1,1,1}, Color: colorful.Color{0,1,0}},
			},
		},
		Spheres: []primitives.Sphere{
			primitives.Sphere{Origin: mgl64.Vec3{0, 0, 0}, Radius: 5.0, Material: materials.Material{Color: colorful.Color{R: 1, G: 0, B: 0}, Shininess: 2}},
		},
	}

	expectedJson := `{"id":8752,"world":{"ambient":{"R":1,"G":0,"B":0},"background":{"R":1,"G":1,"B":1},"camera":{"origin":[10,10,-50],"direction":[0,0,1],"up":[0,1,0],"screenWidth":2,"screenHeight":4,"screenDistance":11},"lights":[{"direction":[1,1,1],"color":{"R":0,"G":1,"B":0}}]},"spheres":[{"origin":[0,0,0],"material":{"color":{"R":1,"G":0,"B":0},"shininess":2},"radius":5}]}`
	testutils.ExpectJsonEncoding(t, &scene, expectedJson)
}

func TestScene_FromPostRequest(t *testing.T) {
	validator := validation.NewValidator()

	testCases := map[string]struct {
		PostRequestJson string
		Expected        model.Scene
	}{
		"empty scene": {
			PostRequestJson: `{"id": 8435}`,
			Expected:        model.Scene{
				Id: 8435,
				World: model.World{
					Ambient:    colorful.Color{0.1, 0.1, 0.1},
					Background: colorful.Color{0, 0, 0},
					Camera:     model.Camera{Origin: mgl64.Vec3{0, 0, -10}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 4, ScreenHeight: 3, ScreenDistance: 1},
					Lights:     []model.Light{},
				},
				Spheres: []primitives.Sphere{},
			},
		},
		"custom scene": {
			PostRequestJson: `{"id":8752,"world":{"ambient":{"R":1,"G":0,"B":0},"background":{"R":1,"G":1,"B":1},"camera":{"origin":[10,10,-50],"direction":[0,0,1],"up":[0,1,0],"screenWidth":2,"screenHeight":4,"screenDistance":11},"lights":[{"direction":[1,1,1],"color":{"R":0,"G":1,"B":0}}]},"spheres":[{"origin":[0,0,0],"material":{"color":{"R":1,"G":0,"B":0},"shininess":2},"radius":5}]}`,
			Expected:        model.Scene{
				Id: 8752,
				World: model.World{
					Ambient:    colorful.Color{1, 0, 0},
					Background: colorful.Color{1, 1, 1},
					Camera:     model.Camera{Origin: mgl64.Vec3{10, 10, -50}, Direction: mgl64.Vec3{0, 0, 1}, Up: mgl64.Vec3{0, 1, 0}, ScreenWidth: 2, ScreenHeight: 4, ScreenDistance: 11},
					Lights:     []model.Light{
						model.Light{Direction: mgl64.Vec3{1,1,1}, Color: colorful.Color{0,1,0}},
					},
				},
				Spheres: []primitives.Sphere{
					primitives.Sphere{Origin: mgl64.Vec3{0, 0, 0}, Radius: 5.0, Material: materials.Material{Color: colorful.Color{R: 1, G: 0, B: 0}, Shininess: 2}},
				},
			},
		},
	}

	validationTestCases := map[string]struct {
		PostRequestJson string
		Expected string
	}{
		"Invalid id": {
			PostRequestJson: `{"id": -3}`,
			Expected: "Id: less than min",
		},
		"Missing id": {
			PostRequestJson: `{}`,
			Expected: "Id: zero value",
		},
		"Zero id": {
			PostRequestJson: `{"id": 0}`,
			Expected: "Id: zero value",
		},
		"Invalid world": {
			PostRequestJson: `{"id": 1, "world": {"ambient": {"R": 0, "G": 0, "B": 2}}}`,
			Expected: "World: not a valid World (Ambient: not a valid color (rgb in 0-1))",
		},
		"Invalid spheres": {
			PostRequestJson: `{"id": 3, "spheres": [{"origin": [1,2,3], "radius": 0}]}`,
			Expected: "Spheres: not a valid []Sphere (Radius: zero value)",
		},
	}

	for name, testCase := range testCases {
		var postRequest model.ScenePostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var scene model.Scene
		if err := scene.FromPostRequest(validator, postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error was returned: %s", name, err.Error())
			continue
		}

		if !scene.Equals(testCase.Expected) {
			t.Errorf("'%s' failed. Expected %#v, received %#v", name, testCase.Expected, scene)
		}
	}

	for name, testCase := range validationTestCases {
		var postRequest model.ScenePostRequest
		if err := json.Unmarshal([]byte(testCase.PostRequestJson), &postRequest); err != nil {
			t.Errorf("'%s' failed. An unexpected error occurred while unmarshaling post request: %s", name, err.Error())
			continue
		}

		var world model.Scene
		if err := world.FromPostRequest(validator, postRequest); err == nil || err.Error() != testCase.Expected {
			t.Errorf("'%s' failed. Expected error %s, received %v", name, testCase.Expected, err)
		}
	}
}

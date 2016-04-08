package phong_test

import (
	"github.com/go-gl/mathgl/mgl64"
	"testing"
	"github.com/bcokert/render-cloud/raytracer/illumination/phong"
	"github.com/lucasb-eyer/go-colorful"
)

func TestCombineColors(t *testing.T) {
	testCases := []struct{
	    C1, C2, Expected colorful.Color
	}{
	    {colorful.Color{0,0,0}, colorful.Color{0,0,0}, colorful.Color{0,0,0}},
	    {colorful.Color{0,1,0}, colorful.Color{0,0,0}, colorful.Color{0,1,0}},
	    {colorful.Color{1,0,1}, colorful.Color{0,1,0}, colorful.Color{1,1,1}},
	    {colorful.Color{0.5,0,1}, colorful.Color{0,0,0}, colorful.Color{0.5,0,1}},
	    {colorful.Color{0.5,0,1}, colorful.Color{0.2,0.2,0}, colorful.Color{0.7,0.2,1}},
	    {colorful.Color{1,1,1}, colorful.Color{1,1,1}, colorful.Color{1,1,1}},
	}

	for i, testCase := range testCases {
	    combined := phong.CombineColors(testCase.C1, testCase.C2)
		if combined.R != testCase.Expected.R || combined.G != testCase.Expected.G || combined.B != testCase.Expected.B {
			t.Errorf("TestCombineColors failed for test case %d. Expected %s, received %s", i, testCase.Expected, combined)
		}
	}
}

func TestSpecular(t *testing.T) {
	testCases := []struct{
	    Reflected mgl64.Vec3
		Viewer mgl64.Vec3
		Shininess float64
		Expected float64
	}{
		// same vector
	    {mgl64.Vec3{1,0,0}, mgl64.Vec3{1,0,0}, 0, 1},
	    {mgl64.Vec3{0,2,0}, mgl64.Vec3{0,5,0}, 1, 1},
	    {mgl64.Vec3{0,0,1}, mgl64.Vec3{0,0,1}, 1, 1},
	    {mgl64.Vec3{1,1,0}, mgl64.Vec3{1,1,0}, 1, 1},
	    {mgl64.Vec3{1,1,1}, mgl64.Vec3{1,1,1}, 1, 1},
	    {mgl64.Vec3{1,1,1}, mgl64.Vec3{1,1,1}, 3, 1},
	    {mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 1, 1},
	    {mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 5, 1},

		// orthogonal
	    {mgl64.Vec3{0,2,0}, mgl64.Vec3{2,0,0}, 1, 0},
	    {mgl64.Vec3{2,0,0}, mgl64.Vec3{0,0,1}, 1, 0},
	    {mgl64.Vec3{0,0,3}, mgl64.Vec3{8,0,0}, 1, 0},

		// real
		{mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 1, 0.8837879163470618},
		{mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 3, 0.6903100211467591},
		{mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 5, 0.5391880975984146},
		{mgl64.Vec3{1,5,5}, mgl64.Vec3{4,11,4}, 1, 0.8943268875682438},
		{mgl64.Vec3{1,5,5}, mgl64.Vec3{4,11,4}, 10, 0.32731271636307313},
	}

	for i, testCase := range testCases {
		if specular := phong.Specular(testCase.Reflected, testCase.Viewer, testCase.Shininess); !mgl64.FloatEqual(specular, testCase.Expected) {
			t.Errorf("TestSpecular for case %d failed. Expected %s, received %s", i, testCase.Expected, specular)
		}
	}
}

func TestDiffuse(t *testing.T) {
	testCases := []struct{
		SpecularComponent float64
		Light mgl64.Vec3
		Normal mgl64.Vec3
		Expected float64
	}{
		// same vector
		{0, mgl64.Vec3{1,0,0}, mgl64.Vec3{1,0,0}, 1},
		{0, mgl64.Vec3{0,2,0}, mgl64.Vec3{0,5,0}, 1},
		{0, mgl64.Vec3{0,0,1}, mgl64.Vec3{0,0,1}, 1},
		{0, mgl64.Vec3{1,1,0}, mgl64.Vec3{1,1,0}, 1},
		{0, mgl64.Vec3{1,1,1}, mgl64.Vec3{1,1,1}, 1},
		{0, mgl64.Vec3{1,1,1}, mgl64.Vec3{1,1,1}, 1},
		{0, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 1},
		{0, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 1},

		// orthogonal
		{0, mgl64.Vec3{0,2,0}, mgl64.Vec3{2,0,0}, 0},
		{0, mgl64.Vec3{2,0,0}, mgl64.Vec3{0,0,1}, 0},
		{0, mgl64.Vec3{0,0,3}, mgl64.Vec3{8,0,0}, 0},

		// real
		{0, mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 0.8837879163470618},
		{0, mgl64.Vec3{1,5,5}, mgl64.Vec3{4,11,4}, 0.8943268875682438},

		// with specular
		{0, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 1},
		{0.1, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 0.9},
		{0.5, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 0.5},
		{0.99, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 0.01},
		{1, mgl64.Vec3{7,2,3}, mgl64.Vec3{7,2,3}, 0},
		{0, mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 0.8837879163470618},
		{0.2, mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 0.70703033307},
		{0.5, mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 0.44189395817},
		{1, mgl64.Vec3{1,6,0}, mgl64.Vec3{-1,3,0}, 0},
	}

	for i, testCase := range testCases {
		if diffuse := phong.Diffuse(testCase.SpecularComponent, testCase.Light, testCase.Normal); !mgl64.FloatEqual(diffuse, testCase.Expected) {
			t.Errorf("TestDiffuse for case %d failed. Expected %s, received %s", i, testCase.Expected, diffuse)
		}
	}
}

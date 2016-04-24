package phong_test

import (
	"github.com/go-gl/mathgl/mgl64"
	"testing"
	"github.com/bcokert/render-cloud/raytracer/illumination/phong"
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/model/materials"
	"strings"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/utils"
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

	phongIlluminator := phong.PhongIlluminator{}
	for i, testCase := range testCases {
	    combined := phongIlluminator.CombineColors(testCase.C1, testCase.C2)
		if !mgl64.FloatEqual(combined.R, testCase.Expected.R) || !mgl64.FloatEqual(combined.G, testCase.Expected.G) || !mgl64.FloatEqual(combined.B, testCase.Expected.B) {
			t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, combined)
		}
	}
}

func TestMultiplyColors(t *testing.T) {
	testCases := []struct{
		C1, C2, Expected colorful.Color
	}{
		{colorful.Color{0,0,0}, colorful.Color{0,0,0}, colorful.Color{0,0,0}},
		{colorful.Color{0,1,0}, colorful.Color{0,0,0}, colorful.Color{0,0,0}},
		{colorful.Color{1,0,1}, colorful.Color{0,1,0}, colorful.Color{0,0,0}},
		{colorful.Color{0.5,0,1}, colorful.Color{0.2,0.2,1}, colorful.Color{0.1,0,1}},
		{colorful.Color{1,1,1}, colorful.Color{1,1,1}, colorful.Color{1,1,1}},
	}

	phongIlluminator := phong.PhongIlluminator{}
	for i, testCase := range testCases {
		combined := phongIlluminator.MultiplyColors(testCase.C1, testCase.C2)
		if !mgl64.FloatEqual(combined.R, testCase.Expected.R) || !mgl64.FloatEqual(combined.G, testCase.Expected.G) || !mgl64.FloatEqual(combined.B, testCase.Expected.B) {
			t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, combined)
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

		// dot product negative
		{mgl64.Vec3{6,1,0}, mgl64.Vec3{-6,1,0}, 1, 0},
		{mgl64.Vec3{0,1,5}, mgl64.Vec3{0,1,-5}, 1, 0},
		{mgl64.Vec3{1,1,1}, mgl64.Vec3{-1,-1,-1}, 1, 0},
	}

	for i, testCase := range testCases {
		specular, err := phong.Specular(testCase.Reflected, testCase.Viewer, testCase.Shininess)
		if err != nil {
			t.Errorf("Case %d failed. Returned error: %s", i, err.Error())
		}
		if !mgl64.FloatEqual(specular, testCase.Expected) {
			t.Errorf("Case %d failed. Expected %#v, received %#v", i, testCase.Expected, specular)
		}
	}
}

func TestSpecularError(t *testing.T) {
	testCases := []struct{
		Reflected mgl64.Vec3
		Viewer mgl64.Vec3
		Shininess float64
		Expected string
	}{
		{mgl64.Vec3{0,0,0}, mgl64.Vec3{1,0,0}, 0, "phong.Specular requires reflectedVector to be a direction vector, received the vector [0, 0, 0]"},
		{mgl64.Vec3{0,1,1}, mgl64.Vec3{0,0,0}, 0, "phong.Specular requires viewerVector to be a direction vector, received the vector [0, 0, 0]"},
	}

	for i, testCase := range testCases {
		_, err := phong.Specular(testCase.Reflected, testCase.Viewer, testCase.Shininess)
		if err == nil || err.Error() != testCase.Expected {
			t.Errorf("Case %d failed. Expected error %s, received %s", i, testCase.Expected, err)
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

		// dot product negative
		{0.5, mgl64.Vec3{6,1,0}, mgl64.Vec3{-6,1,0}, 0},
		{0.5, mgl64.Vec3{0,1,5}, mgl64.Vec3{0,1,-5}, 0},
		{0.5, mgl64.Vec3{1,1,1}, mgl64.Vec3{-1,-1,-1}, 0},
	}

	for i, testCase := range testCases {
		diffuse, err := phong.Diffuse(testCase.SpecularComponent, testCase.Light, testCase.Normal)
		if err != nil {
			t.Errorf("Case %d failed. Returned error: %s", i, err.Error())
		}
		if !mgl64.FloatEqual(diffuse, testCase.Expected) {
			t.Errorf("Case %d failed. Expected %s, received %s", i, testCase.Expected, diffuse)
		}
	}
}

func TestDiffuseError(t *testing.T) {
	testCases := []struct{
		SpecularComponent float64
		Light mgl64.Vec3
		Normal mgl64.Vec3
		Expected string
	}{
		{0, mgl64.Vec3{0,0,0}, mgl64.Vec3{1,0,0}, "phong.Diffuse requires lightVector to be a direction vector, received the vector [0, 0, 0]"},
		{0, mgl64.Vec3{0,1,1}, mgl64.Vec3{0,0,0}, "phong.Diffuse requires normalVector to be a direction vector, received the vector [0, 0, 0]"},
	}

	for i, testCase := range testCases {
		_, err := phong.Diffuse(testCase.SpecularComponent, testCase.Light, testCase.Normal)
		if err == nil || err.Error() != testCase.Expected {
			t.Errorf("Case %d failed. Expected error %s, received %s", i, testCase.Expected, err)
		}
	}
}

func TestViewerVector(t *testing.T) {
	testCases := []struct{
	    Ray mgl64.Vec3
		Expected mgl64.Vec3
	}{
	    {mgl64.Vec3{0,1,0}, mgl64.Vec3{0,-1,0}},
	    {mgl64.Vec3{1,1,1}, mgl64.Vec3{-1,-1,-1}.Normalize()},
	    {mgl64.Vec3{9,-1,4}, mgl64.Vec3{-9,1,-4}.Normalize()},
	}

	for i, testCase := range testCases {
		result, err := phong.ViewerVector(testCase.Ray)
		if err != nil {
			t.Errorf("Case %d failed. Unexpected error: %s", i, err.Error())
		}
		if !result.ApproxEqual(testCase.Expected) {
			t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, result)
		}
	}
}

func TestViewerVectorError(t *testing.T) {
	testCases := []struct{
	    Viewer mgl64.Vec3
	    Expected string
	}{
	    {mgl64.Vec3{0,0,0}, "phong.ViewerVector requires ray to be a direction vector, received the vector [0, 0, 0]"},
	}

	for i, testCase := range testCases {
	    _, err := phong.ViewerVector(testCase.Viewer)
	    if err == nil {
	        t.Errorf("Case %d failed. Expected error but none occured", i)
	    }
	    if err.Error() != testCase.Expected {
	        t.Errorf("Case $d failed. Expected %s, received %s", i, testCase.Expected, err.Error())
	    }
	}
}

func TestLightVector(t *testing.T) {
	testCases := []struct{
		Light model.Camera
		Expected mgl64.Vec3
	}{
		{model.Camera{Direction: &mgl64.Vec3{0,1,0}}, mgl64.Vec3{0,-1,0}},
		{model.Camera{Direction: &mgl64.Vec3{1,1,1}}, mgl64.Vec3{-1,-1,-1}.Normalize()},
		{model.Camera{Direction: &mgl64.Vec3{9,-1,4}}, mgl64.Vec3{-9,1,-4}.Normalize()},
	}

	for i, testCase := range testCases {
		result, err := phong.LightVector(testCase.Light)
		if err != nil {
			t.Errorf("Case %d failed. Unexpected error: %s", i, err.Error())
		}
		if !result.ApproxEqual(testCase.Expected) {
			t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, result)
		}
	}
}

func TestLightVectorError(t *testing.T) {
	testCases := []struct{
		Light model.Camera
	    Expected string
	}{
	    {model.Camera{Direction: &mgl64.Vec3{0,0,0}}, "phong.LightVector requires light.Direction to be a direction vector, received the vector [0, 0, 0]"},
	}

	for i, testCase := range testCases {
	    _, err := phong.LightVector(testCase.Light)
	    if err == nil {
	        t.Errorf("Case %d failed. Expected error but none occured", i)
	    }
	    if err.Error() != testCase.Expected {
	        t.Errorf("Case %d failed. Expected %s, received %s", i, testCase.Expected, err.Error())
	    }
	}
}

func TestReflectedVector(t *testing.T) {
	testCases := []struct{
		Normal mgl64.Vec3
		Light mgl64.Vec3
	    Expected mgl64.Vec3
	}{
	    {mgl64.Vec3{0,1,0}, mgl64.Vec3{1,1,0}, mgl64.Vec3{-1,1,0}.Normalize()},
	    {mgl64.Vec3{0,4,0}, mgl64.Vec3{2,2,0}, mgl64.Vec3{-2,2,0}.Normalize()},
	    {mgl64.Vec3{0,0,2}, mgl64.Vec3{2,0,2}, mgl64.Vec3{-2,0,2}.Normalize()},
	    {mgl64.Vec3{1,1,1}, mgl64.Vec3{3,3,3}, mgl64.Vec3{3,3,3}.Normalize()},
	    {mgl64.Vec3{1,1,1}, mgl64.Vec3{3,3,3}, mgl64.Vec3{3,3,3}.Normalize()},
	}

	for i, testCase := range testCases {
	    result, err := phong.ReflectedVector(testCase.Normal, testCase.Light)
	    if err != nil {
	        t.Errorf("Case %d failed. Unexpected error %s", i, err.Error())
	    }
	    if !result.ApproxEqual(testCase.Expected) {
	        t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, result)
	    }
	}
}

func TestReflectedVectorError(t *testing.T) {
	testCases := []struct{
	    Normal mgl64.Vec3
		Light mgl64.Vec3
	    Expected string
	}{
	    {mgl64.Vec3{0,0,0}, mgl64.Vec3{1,1,1}, "phong.LightVector requires normalVector to be a direction vector, received the vector [0, 0, 0]"},
	    {mgl64.Vec3{1,1,1}, mgl64.Vec3{0,0,0}, "phong.LightVector requires lightVector to be a direction vector, received the vector [0, 0, 0]"},
	}

	for i, testCase := range testCases {
	    _, err := phong.ReflectedVector(testCase.Normal, testCase.Light)
	    if err == nil {
	        t.Errorf("Case %d failed. Expected error but none occured", i)
	    }
	    if err.Error() != testCase.Expected {
	        t.Errorf("Case %d failed. Expected %s, received %s", i, testCase.Expected, err.Error())
	    }
	}
}

func TestIlluminateLocal(t *testing.T) {
	phongIlluminator := phong.PhongIlluminator{}
	testCases := []struct{
		Ray mgl64.Vec3
		Normal mgl64.Vec3
		Material materials.Material
		World model.World
		Expected colorful.Color
	}{
		// No Lights
		{Ray: mgl64.Vec3{1,1,0}, Normal: mgl64.Vec3{0,1,0}, Material: materials.Material{Color: &colorful.Color{1,0,0}, Shininess: utils.FloatPointer(1)}, World: model.World{Ambient: &colorful.Color{0.5,0.5,0.5}, Lights: &[]model.Camera{}}, Expected: colorful.Color{0.5,0,0}},
		{Ray: mgl64.Vec3{1,1,0}, Normal: mgl64.Vec3{0,1,0}, Material: materials.Material{Color: &colorful.Color{0,0.5,1}, Shininess: utils.FloatPointer(1)}, World: model.World{Ambient: &colorful.Color{0.5,0.5,0.5}, Lights: &[]model.Camera{}}, Expected: colorful.Color{0,0.25,0.5}},
		{Ray: mgl64.Vec3{1,1,0}, Normal: mgl64.Vec3{0,1,0}, Material: materials.Material{Color: &colorful.Color{1,1,1}, Shininess: utils.FloatPointer(1)}, World: model.World{Ambient: &colorful.Color{0.3,0.5,0.7}, Lights: &[]model.Camera{}}, Expected: colorful.Color{0.3,0.5,0.7}},

		// Specular + Ambient
		{mgl64.Vec3{0,0,10}, mgl64.Vec3{0,0,3}, materials.Material{Color: &colorful.Color{0.5,0.5,0.5}, Shininess: utils.FloatPointer(1)}, model.World{Ambient: &colorful.Color{0.1,0.1,0.1}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}, Color: &colorful.Color{0.3,0.3,0.3}}}}, phongIlluminator.CombineColors(colorful.Color{0.05,0.05,0.05}, colorful.Color{0.3,0.3,0.3})},
		{mgl64.Vec3{0,0,10}, mgl64.Vec3{0,0,3}, materials.Material{Color: &colorful.Color{1,0,0}, Shininess: utils.FloatPointer(5)}, model.World{Ambient: &colorful.Color{0.2,0.2,0.2}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}, Color: &colorful.Color{0.3,0.4,0.5}}}}, phongIlluminator.CombineColors(colorful.Color{0.2,0,0}, colorful.Color{0.3,0.4,0.5})},

		// Specular + Diffuse + Ambient
		{mgl64.Vec3{2,0,10}, mgl64.Vec3{0,0,1}, materials.Material{Color: &colorful.Color{0.2,0.2,0.8}, Shininess: utils.FloatPointer(1)}, model.World{Ambient: &colorful.Color{0.2,0.2,0.2}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}, Color: &colorful.Color{0.6,0.6,0.6}}}}, colorful.Color{0.6283484054145522, 0.6283484054145522, 0.7483484054145522}},
		{mgl64.Vec3{2,0,10}, mgl64.Vec3{0,0,1}, materials.Material{Color: &colorful.Color{0.2,0.2,0.8}, Shininess: utils.FloatPointer(5)}, model.World{Ambient: &colorful.Color{0.2,0.2,0.2}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}, Color: &colorful.Color{0.6,0.6,0.6}}}}, colorful.Color{0.5839611736451112, 0.5839611736451112, 0.7039611736451112}},
		{mgl64.Vec3{2,0,10}, mgl64.Vec3{0,0,1}, materials.Material{Color: &colorful.Color{0.2,0.2,0.8}, Shininess: utils.FloatPointer(1)}, model.World{Ambient: &colorful.Color{0.2,0.2,0.2}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}, Color: &colorful.Color{0.6,0,0}}}}, colorful.Color{0.6283484054145522, 0.04, 0.16}},
		{mgl64.Vec3{2,0,10}, mgl64.Vec3{0,0,1}, materials.Material{Color: &colorful.Color{0.2,0.2,0.8}, Shininess: utils.FloatPointer(5)}, model.World{Ambient: &colorful.Color{0.2,0.2,0.2}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}, Color: &colorful.Color{0.6,0,0}}}}, colorful.Color{0.5839611736451112, 0.04, 0.16}},
	}

	for i, testCase := range testCases {
		result, err := phongIlluminator.IlluminateLocal(testCase.Ray, testCase.Normal, testCase.Material, testCase.World)
		if err != nil {
			t.Errorf("Case %d failed. Returned unexpected error: %s", i, err.Error())
		} else {
			if !utils.ColorsApproxEqual(testCase.Expected, result) {
				t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, result)
			}
		}
	}
}


func TestIlluminateLocalError(t *testing.T) {
	testCases := []struct{
		Ray mgl64.Vec3
		Normal mgl64.Vec3
		Material materials.Material
		World model.World
		Expected string
	}{
		{mgl64.Vec3{0,0,0}, mgl64.Vec3{0,0,1}, materials.Material{Color: &colorful.Color{0,0,0}}, model.World{Ambient: &colorful.Color{0,0,0}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}}}}, "phong.ViewerVector Failed"},
		{mgl64.Vec3{0,0,1}, mgl64.Vec3{0,0,0}, materials.Material{Color: &colorful.Color{0,0,0}}, model.World{Ambient: &colorful.Color{0,0,0}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,1}}}}, "phong.ReflectedVector Failed"},
		{mgl64.Vec3{0,0,1}, mgl64.Vec3{0,0,1}, materials.Material{Color: &colorful.Color{0,0,0}}, model.World{Ambient: &colorful.Color{0,0,0}, Lights: &[]model.Camera{model.Camera{Direction: &mgl64.Vec3{0,0,0}}}}, "phong.LightVector Failed"},
	}

	phongIlluminator := phong.PhongIlluminator{}
	for i, testCase := range testCases {
		_, err := phongIlluminator.IlluminateLocal(testCase.Ray, testCase.Normal, testCase.Material, testCase.World)
		if err == nil {
			t.Errorf("Case %d failed. An error was not thrown when illegal input was given", i)
		} else {
			if !strings.HasPrefix(err.Error(), testCase.Expected) {
				t.Errorf("Case %d failed. Expected error %s, received error %s", i, testCase.Expected, err.Error())
			}
		}
	}
}

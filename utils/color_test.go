package utils_test

import (
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/bcokert/render-cloud/utils"
)

func TestScaleColor(t *testing.T) {
	testCases := []struct{
		Color colorful.Color
		Factor float64
		ExpectedColor colorful.Color
	}{
	    {colorful.Color{0,0,0}, 0, colorful.Color{0,0,0}},
	    {colorful.Color{0,0,0}, 5, colorful.Color{0,0,0}},
	    {colorful.Color{1,1,1}, 5, colorful.Color{1,1,1}},
	    {colorful.Color{1,1,1}, 0.1, colorful.Color{0.1,0.1,0.1}},
	    {colorful.Color{0,0.5,1}, 0.5, colorful.Color{0,0.25,0.5}},
	    {colorful.Color{0.1, 0.5, 0.9}, 0.23, colorful.Color{0.023,0.115,0.207}},
	    {colorful.Color{0, 0, 1}, 0.5, colorful.Color{0,0,0.5}},
	    {colorful.Color{0, 1, 0}, 0.5, colorful.Color{0,0.5,0}},
	    {colorful.Color{1, 0, 0}, 0.5, colorful.Color{0.5,0,0}},
	}

	for i, testCase := range testCases {
		scaledColor := utils.ScaleColor(testCase.Color, testCase.Factor)
		if !scaledColor.AlmostEqualRgb(testCase.ExpectedColor) {
			t.Errorf("Scaling testcase %d was incorrect. Expected %s, found %s.", i, testCase.ExpectedColor, scaledColor)
		}
	}
}

package validation_test

import ()
import (
	"testing"
	"github.com/bcokert/render-cloud/validation"
	"github.com/go-gl/mathgl/mgl64"
	"errors"
	"github.com/lucasb-eyer/go-colorful"
)

func TestNewValidator_implements(t *testing.T) {
	var _ validation.Validator = validation.NewValidator()
}

type DirectionTestStruct struct {
	Dir interface{} `validate:"direction"`
}

func TestNewValidator_direction(t *testing.T) {
	val := validation.NewValidator()

    testCases := map[string]struct{
	    Input interface{}
        Expected error
    }{
        "no zeros": {
	        Input: DirectionTestStruct{mgl64.Vec3{1,2,-3}},
	        Expected: nil,
        },
	    "some zeros": {
		    Input: DirectionTestStruct{mgl64.Vec3{0,1,0}},
		    Expected: nil,
	    },
	    "all zeros": {
	        Input: DirectionTestStruct{mgl64.Vec3{0,0,0}},
		    Expected: errors.New("Dir: not a direction vector (found [0,0,0])"),
	    },
	    "wrong type" : {
		    Input: DirectionTestStruct{32},
		    Expected: errors.New("Dir: not a vector type"),
	    },
    }

    for name, testCase := range testCases {
	    result := val.Validate(testCase.Input)
	    if (testCase.Expected == nil && result != nil) || (testCase.Expected != nil && result == nil) || (testCase.Expected != nil && testCase.Expected.Error() != result.Error()) {
            t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
        }
    }
}

type ColorTestStruct struct {
	Col interface{} `validate:"color"`
}

func TestNewValidator_color(t *testing.T) {
	val := validation.NewValidator()

	testCases := map[string]struct{
		Input interface{}
		Expected error
	}{
		"all ones": {
			Input: ColorTestStruct{colorful.Color{1,1,1}},
			Expected: nil,
		},
		"valid color": {
			Input: ColorTestStruct{colorful.Color{0,0.2,0.7}},
			Expected: nil,
		},
		"all zeros": {
			Input: ColorTestStruct{colorful.Color{0,0,0}},
			Expected: nil,
		},
		"one bad": {
			Input: ColorTestStruct{colorful.Color{0,-0.1,0}},
			Expected: errors.New("Col: not a valid color (rgb in 0-1)"),
		},
		"all bad": {
			Input: ColorTestStruct{colorful.Color{1.1,1.6,2.3}},
			Expected: errors.New("Col: not a valid color (rgb in 0-1)"),
		},
		"wrong type": {
			Input: ColorTestStruct{mgl64.Vec3{1.1,1.6,2.3}},
			Expected: errors.New("Col: not a color type"),
		},
	}

	for name, testCase := range testCases {
		result := val.Validate(testCase.Input)
		if (testCase.Expected == nil && result != nil) || (testCase.Expected != nil && result == nil) || (testCase.Expected != nil && testCase.Expected.Error() != result.Error()) {
			t.Errorf("'%s' failed. Expected %v, received %v", name, testCase.Expected, result)
		}
	}
}

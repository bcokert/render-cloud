package validation

import (
	"errors"
	"github.com/go-gl/mathgl/mgl64"
	"gopkg.in/validator.v2"
	"github.com/lucasb-eyer/go-colorful"
)

// A Validator can validate any given struct by using its struct tags
// Internally uses gopkg.in/validator.v2.Validator by default
type Validator interface {
	Validate(v interface{}) error
}

// Returns a new Validator that contains all custom validators
func NewValidator () Validator {
	v := validator.NewValidator()
	v.SetValidationFunc("direction", direction)
	v.SetValidationFunc("color", color)

	return Validator(v)
}

// Validates that a given interface is a direction vector
// All vectors expect [0,0,0] have a direction
func direction(v interface{}, param string) error {
	t, ok := v.(mgl64.Vec3)
	if !ok {
		return errors.New("not a vector type")
	}

	invalid := mgl64.Vec3{0,0,0}
	if t == invalid {
		return errors.New("not a direction vector (found [0,0,0])")
	}

	return nil
}

// Validates that a given interface is a valid color
// All colors have R,G,B components between 0 and 1, inclusive
func color(v interface{}, param string) error {
	t, ok := v.(colorful.Color)
	if !ok {
		return errors.New("not a color type")
	}

	if !t.IsValid() {
		return errors.New("not a valid color (rgb in 0-1)")
	}

	return nil
}

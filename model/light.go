package model

import (
	"github.com/bcokert/render-cloud/utils"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/validation"
)

type LightPostRequest struct {
	Direction *mgl64.Vec3     `json:"direction"`
	Color     *colorful.Color `json:"color"`
}

type Light struct {
	Direction mgl64.Vec3          `json:"direction" validate:"direction"`
	Color     colorful.Color      `json:"color" validate:"color"`
}

func (this *Light) FromPostRequest(validator validation.Validator, postRequest LightPostRequest) error {
	if err := validator.Validate(postRequest); err != nil {
		return err
	}

	this.Direction = utils.DefaultVector(postRequest.Direction, mgl64.Vec3{-1, -1, 1})
	this.Color = utils.DefaultColor(postRequest.Color, colorful.Color{0.7, 0.7, 0.7})

	return validator.Validate(this)
}

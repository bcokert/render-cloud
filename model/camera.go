package model

import (
	"github.com/bcokert/render-cloud/utils"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/bcokert/render-cloud/validation"
)

type CameraPostRequest struct {
	Origin         *mgl64.Vec3         `json:"origin"`
	Direction      *mgl64.Vec3         `json:"direction"`
	Up             *mgl64.Vec3         `json:"up"`
	ScreenWidth    *float64            `json:"screenWidth"`
	ScreenHeight   *float64            `json:"screenHeight"`
	ScreenDistance *float64            `json:"screenDistance"`
}

type Camera struct {
	Origin         mgl64.Vec3          `json:"origin"`
	Direction      mgl64.Vec3          `json:"direction" validate:"direction"`
	Up             mgl64.Vec3          `json:"up" validate:"direction"`
	ScreenWidth    float64             `json:"screenWidth" validate:"nonzero,min=0"`
	ScreenHeight   float64             `json:"screenHeight" validate:"nonzero,min=0"`
	ScreenDistance float64             `json:"screenDistance" validate:"nonzero,min=0"`
}

func (this *Camera) FromPostRequest(validator validation.Validator, postRequest CameraPostRequest) error {
	if err := validator.Validate(postRequest); err != nil {
		return err
	}

	this.Origin = utils.DefaultVector(postRequest.Origin, mgl64.Vec3{0, 0, -10})
	this.Direction = utils.DefaultVector(postRequest.Direction, mgl64.Vec3{0, 0, 1})
	this.Up = utils.DefaultVector(postRequest.Up, mgl64.Vec3{0, 1, 0})
	this.ScreenWidth = utils.DefaultFloat(postRequest.ScreenWidth, 4)
	this.ScreenHeight = utils.DefaultFloat(postRequest.ScreenHeight, 3)
	this.ScreenDistance = utils.DefaultFloat(postRequest.ScreenDistance, 1)

	return validator.Validate(this)
}

package model

import (
	"github.com/bcokert/render-cloud/utils"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/validation"
	"fmt"
)

type WorldPostRequest struct {
	Ambient    *colorful.Color   `json:"ambient"`
	Background *colorful.Color   `json:"background"`
	Camera     CameraPostRequest `json:"camera"`
	Lights     []LightPostRequest `json:"lights"`
}

type World struct {
	Ambient    colorful.Color       `json:"ambient" validate:"color"`
	Background colorful.Color       `json:"background" validate:"color"`
	Camera     Camera               `json:"camera"`
	Lights     []Light               `json:"lights"`
}

func (this World) Equals(world World) bool {
	if this.Ambient != world.Ambient || this.Background != world.Background || this.Camera != world.Camera {
		return false
	}

	if len(this.Lights) != len(world.Lights) {
		return false
	}

	for i, light := range world.Lights {
		if light != this.Lights[i] {
			return false
		}
	}

	return true
}

func (this *World) FromPostRequest(validator validation.Validator, postRequest WorldPostRequest) error {
	if err := validator.Validate(postRequest); err != nil {
		return err
	}

	this.Ambient = utils.DefaultColor(postRequest.Ambient, colorful.Color{0.1, 0.1, 0.1})
	this.Background = utils.DefaultColor(postRequest.Background, colorful.Color{0, 0, 0})
	if err := this.Camera.FromPostRequest(validator, postRequest.Camera); err != nil {
		return fmt.Errorf("Camera: not a valid Camera (%s)", err.Error())
	}
	this.Lights = []Light{}
	for _, lightRequest := range postRequest.Lights {
		var light Light
		if err := light.FromPostRequest(validator, lightRequest); err != nil {
			return fmt.Errorf("Lights: not a valid []Light (%s)", err.Error())
		}

		this.Lights = append(this.Lights, light)
	}

	return validator.Validate(this)
}

package materials

import (
	"github.com/bcokert/render-cloud/utils"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/validation"
)

type MaterialPostRequest struct {
	Color     *colorful.Color `json:"color"`
	Shininess *float64        `json:"shininess"`
}

type Material struct {
	Color     colorful.Color       `json:"color" validate:"color"`
	Shininess float64              `json:"shininess" validate:"min=0,nonzero"`
}

func (this *Material) FromPostRequest(validator validation.Validator, postRequest MaterialPostRequest) error {
	if err := validator.Validate(postRequest); err != nil {
		return err
	}

	this.Color = utils.DefaultColor(postRequest.Color, colorful.Color{0.8, 0.1, 0.1})
	this.Shininess = utils.DefaultFloat(postRequest.Shininess, 1)

	return validator.Validate(this)
}

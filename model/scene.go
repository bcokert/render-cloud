package model

import (
	"github.com/bcokert/render-cloud/model/primitives"
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
	"io"
)

type ScenePostRequest struct {
	Id      *uint               `json:"id" validate:"nonzero"`
	World   World               `json:"world"`
	Spheres []primitives.Sphere `json:"spheres"`
}

type Scene struct {
	Id      uint                `json:"id" validate:"nonzero"`
	World   World               `json:"world"`
	Spheres []primitives.Sphere `json:"spheres"`
}

func UnmarshalScene(jsonData io.Reader) (*Scene, error) {
	var postRequest ScenePostRequest
	if err := json.NewDecoder(jsonData).Decode(&postRequest); err != nil {
		return nil, fmt.Errorf("UnmarshalScene failed to decode post request: %s", err.Error())
	}

	if errs := validator.Validate(postRequest); errs != nil {
		return nil, fmt.Errorf("UnmarshalScene failed to validate post request: %s", errs.Error())
	}

	scene := Scene{
		*postRequest.Id,
		postRequest.World,
		postRequest.Spheres,
	}

	if errs := validator.Validate(scene); errs != nil {
		return nil, fmt.Errorf("UnmarshalScene failed to validate model: %s", errs.Error())
	}

	return &scene, nil
}

func (this Scene) Encode(writer io.Writer) error {
	if err := json.NewEncoder(writer).Encode(this); err != nil {
		return fmt.Errorf("Marshal failed to encode scene: %s", err.Error())
	}
	return nil
}

package model

import (
	"fmt"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/bcokert/render-cloud/validation"
)

type ScenePostRequest struct {
	Id      *int                           `json:"id" validate:"nonzero"`
	World   WorldPostRequest               `json:"world"`
	Spheres []primitives.SpherePostRequest `json:"spheres"`
}

type Scene struct {
	Id      int                 `json:"id" validate:"min=1"`
	World   World               `json:"world"`
	Spheres []primitives.Sphere `json:"spheres"`
}

func (this Scene) Equals(scene Scene) bool {
	if this.Id != scene.Id || !this.World.Equals(scene.World) {
		return false
	}

	if len(this.Spheres) != len(scene.Spheres) {
		return false
	}

	for i, sphere := range scene.Spheres {
		if sphere != this.Spheres[i] {
			return false
		}
	}

	return true
}

func (this *Scene) FromPostRequest(validator validation.Validator, postRequest ScenePostRequest) error {
	if err := validator.Validate(postRequest); err != nil {
		return err
	}

	this.Id = *postRequest.Id
	if err := this.World.FromPostRequest(validator, postRequest.World); err != nil {
		return fmt.Errorf("World: not a valid World (%s)", err.Error())
	}
	this.Spheres = []primitives.Sphere{}
	for _, sphereRequest := range postRequest.Spheres {
		var sphere primitives.Sphere
		if err := sphere.FromPostRequest(validator, sphereRequest); err != nil {
			return fmt.Errorf("Spheres: not a valid []Sphere (%s)", err.Error())
		}

		this.Spheres = append(this.Spheres, sphere)
	}

	return validator.Validate(this)
}

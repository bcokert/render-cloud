package model

import "github.com/bcokert/render-cloud/model/primitives"

type Scene struct {
	Id      *uint    `json:"id,omitempty"`
	World   *World   `json:"world,omitempty"`
	Spheres *[]primitives.Sphere `json:"spheres,omitempty"`
}

func (scene Scene) GetId() uint {
	return *scene.Id
}

func (scene Scene) GetWorld() World {
	return *scene.World
}

func (scene Scene) GetSpheres() []primitives.Sphere {
	return *scene.Spheres
}

package model

import "github.com/bcokert/render-cloud/model/primitives"

type Scene struct {
	Id      *uint    `json:"id,omitempty"`
	World   *World   `json:"world,omitempty"`
	Spheres *[]primitives.Sphere `json:"spheres,omitempty"`
}

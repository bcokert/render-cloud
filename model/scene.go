package model

type Scene struct {
	Id      *uint    `json:"id,omitempty"`
	World   *World   `json:"world,omitempty"`
	Spheres *Spheres `json:"spheres,omitempty"`
}

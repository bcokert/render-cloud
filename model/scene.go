package model

type Scene struct {
	Id      uint    `json:"id"`
	World   World   `json:"world"`
	Spheres Spheres `json:"spheres"`
}

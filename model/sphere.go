package model

import (
	"github.com/bcokert/render-cloud/vector"
	"github.com/lucasb-eyer/go-colorful"
)

type Sphere struct {
	Origin    *vector.Vector3 `json:"origin,omitempty"`
	Radius    *float64        `json:"radius,omitempty"`
	Color     *colorful.Color `json:"color,omitempty"`
	Shininess *float64        `json:"shininess,omitempty"`
}

type Spheres []Sphere

package model

import (
	"github.com/bcokert/render-cloud/vector"
	"github.com/lucasb-eyer/go-colorful"
)

type Sphere struct {
	Origin    vector.Vector3 `json:"origin"`
	Radius    float64        `json:"radius"`
	Color     colorful.Color `json:"color"`
	Shininess float64        `json:"shininess"`
}

type Spheres []Sphere

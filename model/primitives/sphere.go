package primitives

import (
	"github.com/bcokert/render-cloud/vector"
	"github.com/bcokert/render-cloud/model/materials"
)

type Sphere struct {
	Origin    *vector.Vector3 `json:"origin,omitempty"`
	Radius    *float64        `json:"radius,omitempty"`
	Material  *materials.Material `json:"material,omitempty"`
}

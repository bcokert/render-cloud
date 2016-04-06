package model

import (
	"github.com/bcokert/render-cloud/vector"
)

type Camera struct {
	Origin    *vector.Vector3 `json:"origin,omitempty"`
	Direction *vector.Vector3 `json:"direction,omitempty"`
	Up        *vector.Vector3 `json:"up,omitempty"`
}

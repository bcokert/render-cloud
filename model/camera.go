package model

import (
	"github.com/bcokert/render-cloud/vector"
)

type Camera struct {
	Origin    vector.Vector3 `json:"origin"`
	Direction vector.Vector3 `json:"direction"`
	Up        vector.Vector3 `json:"up"`
}

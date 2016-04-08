package model

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Camera struct {
	Origin    *mgl64.Vec3 `json:"origin,omitempty"`
	Direction *mgl64.Vec3 `json:"direction,omitempty"`
	Up        *mgl64.Vec3 `json:"up,omitempty"`
}

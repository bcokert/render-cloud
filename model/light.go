package model

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
)

type Light struct {
	Direction *mgl64.Vec3     `json:"direction,omitempty"`
	Color     *colorful.Color `json:"color,omitempty"`
}

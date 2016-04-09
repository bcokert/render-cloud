package model

import (
	"github.com/go-gl/mathgl/mgl64"
)

type Camera struct {
	Origin         *mgl64.Vec3 `json:"origin,omitempty"`
	Direction      *mgl64.Vec3 `json:"direction,omitempty"`
	Up             *mgl64.Vec3 `json:"up,omitempty"`
	ScreenWidth    *float64    `json:"screenWidth,omitempty"`
	ScreenHeight   *float64    `json:"screenHeight,omitempty"`
	ScreenDistance *float64    `json:"screenDistance,omitempty"`
}

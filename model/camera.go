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

func (camera Camera) GetOrigin() mgl64.Vec3 {
	return *camera.Origin
}

func (camera Camera) GetDirection() mgl64.Vec3 {
	return *camera.Direction
}

func (camera Camera) GetUp() mgl64.Vec3 {
	return *camera.Up
}

func (camera Camera) GetScreenWidth() float64 {
	return *camera.ScreenWidth
}

func (camera Camera) GetScreenHeight() float64 {
	return *camera.ScreenHeight
}

func (camera Camera) GetScreenDistance() float64 {
	return *camera.ScreenDistance
}

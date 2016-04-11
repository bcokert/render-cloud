package model

import (
	"github.com/lucasb-eyer/go-colorful"
)

type World struct {
	Ambient    *colorful.Color `json:"ambient,omitempty"`
	Background *colorful.Color `json:"background,omitempty"`
	Camera     *Camera      `json:"camera,omitempty"`
	Lights     *[]Light     `json:"lights,omitempty"`
}

func (world World) GetAmbient() colorful.Color {
	return *world.Ambient
}

func (world World) GetBackground() colorful.Color {
	return *world.Background
}

func (world World) GetCamera() Camera {
	return *world.Camera
}

func (world World) GetLights() []Light {
	return *world.Lights
}

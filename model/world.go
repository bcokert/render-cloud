package model

import (
	color "github.com/lucasb-eyer/go-colorful"
)

type World struct {
	Ambient    *color.Color `json:"ambient,omitempty"`
	Background *color.Color `json:"background,omitempty"`
	Camera     *Camera      `json:"camera,omitempty"`
	Lights     *[]Light     `json:"lights,omitempty"`
}

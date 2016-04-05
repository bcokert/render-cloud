package model

import (
	color "github.com/lucasb-eyer/go-colorful"
)

type World struct {
	Ambient    color.Color `json:"ambient"`
	Background color.Color `json:"background"`
	Camera     Camera      `json:"camera"`
}

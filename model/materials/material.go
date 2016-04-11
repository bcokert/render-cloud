package materials

import "github.com/lucasb-eyer/go-colorful"

type Material struct {
	Color     *colorful.Color `json:"color,omitempty"`
	Shininess *float64        `json:"shininess,omitempty"`
}

func (material Material) GetColor() colorful.Color {
	return *material.Color
}

func (material Material) GetShininess() float64 {
	return *material.Shininess
}

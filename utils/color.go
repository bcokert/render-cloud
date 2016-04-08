package utils

import "github.com/lucasb-eyer/go-colorful"

func ScaleColor(color colorful.Color, factor float64) colorful.Color {
	return colorful.Color{color.R * factor, color.G * factor, color.B * factor}.Clamped()
}

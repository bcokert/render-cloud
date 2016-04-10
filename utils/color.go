package utils

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/go-gl/mathgl/mgl64"
)

func ScaleColor(color colorful.Color, factor float64) colorful.Color {
	return colorful.Color{color.R * factor, color.G * factor, color.B * factor}.Clamped()
}

func ColorsApproxEqual(c1, c2 colorful.Color) bool {
	v1 := mgl64.Vec3{c1.R, c1.G, c1.B}
	v2 := mgl64.Vec3{c2.R, c2.G, c2.B}
	return v1.ApproxEqual(v2)
}

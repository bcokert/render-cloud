package utils

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/go-gl/mathgl/mgl64"
)

func DefaultColor(maybe *colorful.Color, def colorful.Color) colorful.Color {
	if maybe != nil {
		return *maybe
	}

	return def
}

func DefaultVector(maybe *mgl64.Vec3, def mgl64.Vec3) mgl64.Vec3 {
	if maybe != nil {
		return *maybe
	}

	return def
}

func DefaultFloat(maybe *float64, def float64) float64 {
	if maybe != nil {
		return *maybe
	}

	return def
}

package phong

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func Specular(reflectedVector, viewerVector mgl64.Vec3, shininess float64) float64 {
	return math.Pow(reflectedVector.Normalize().Dot(viewerVector.Normalize()), shininess)
}

func Diffuse(specularComponent float64, lightVector, normalVector mgl64.Vec3) float64 {
	return (1 - specularComponent) * lightVector.Normalize().Dot(normalVector.Normalize())
}

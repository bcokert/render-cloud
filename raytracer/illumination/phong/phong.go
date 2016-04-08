package phong

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"github.com/lucasb-eyer/go-colorful"
)

func CombineColors(c1, c2 colorful.Color) colorful.Color {
	return colorful.Color{
		c1.R + c2.R,
		c1.G + c2.G,
		c1.B + c2.B,
	}.Clamped()
}

func Specular(reflectedVector, viewerVector mgl64.Vec3, shininess float64) float64 {
	return math.Pow(reflectedVector.Normalize().Dot(viewerVector.Normalize()), shininess)
}

func Diffuse(specularComponent float64, lightVector, normalVector mgl64.Vec3) float64 {
	return (1 - specularComponent) * lightVector.Normalize().Dot(normalVector.Normalize())
}

func IlluminateLocal(ambientColor, specularColor, diffuseColor colorful.Color, lightVector, normalVector, viewerVector, reflectedVector mgl64.Vec3, shininess float64) {
	//specularComponent := Specular(reflectedVector, viewerVector, shininess)
	//diffuseComponent := Diffuse(specularComponent, lightVector, normalVector)
	//ambientComponent := 1.0
	//
	//ambientContributionColor := utils.ScaleColor(ambientColor, ambientComponent)
	//specularContributionColor := utils.ScaleColor(specularColor, specularComponent)
	//diffuseContributionColor := utils.ScaleColor(diffuseColor, diffuseComponent)
}

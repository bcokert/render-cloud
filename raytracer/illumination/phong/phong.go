package phong

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/utils"
	"errors"
)

func CombineColors(c1, c2 colorful.Color) colorful.Color {
	return colorful.Color{
		c1.R + c2.R,
		c1.G + c2.G,
		c1.B + c2.B,
	}.Clamped()
}

func Specular(reflectedVector, viewerVector mgl64.Vec3, shininess float64) (float64, error) {
	if reflectedVector.X() == 0 && reflectedVector.Y() == 0 && reflectedVector.Z() == 0 {
		return 0, errors.New("phong.Specular requires reflectedVector to be a direction vector, received the vector [0, 0, 0]")
	}
	if viewerVector.X() == 0 && viewerVector.Y() == 0 && viewerVector.Z() == 0 {
		return 0, errors.New("phong.Specular requires viewerVector to be a direction vector, received the vector [0, 0, 0]")
	}
	return math.Pow(reflectedVector.Normalize().Dot(viewerVector.Normalize()), shininess), nil
}

func Diffuse(specularComponent float64, lightVector, normalVector mgl64.Vec3) (float64, error) {
	if lightVector.X() == 0 && lightVector.Y() == 0 && lightVector.Z() == 0 {
		return 0, errors.New("phong.Diffuse requires lightVector to be a direction vector, received the vector [0, 0, 0]")
	}
	if normalVector.X() == 0 && normalVector.Y() == 0 && normalVector.Z() == 0 {
		return 0, errors.New("phong.Diffuse requires normalVector to be a direction vector, received the vector [0, 0, 0]")
	}
	return (1 - specularComponent) * lightVector.Normalize().Dot(normalVector.Normalize()), nil
}

func IlluminateLocal(ambientColor, specularColor, diffuseColor colorful.Color, lightVector, normalVector, viewerVector, reflectedVector mgl64.Vec3, shininess float64) (colorful.Color, error) {
	specularComponent, errSpec := Specular(reflectedVector, viewerVector, shininess)
	diffuseComponent, errDiff := Diffuse(specularComponent, lightVector, normalVector)
	ambientComponent := 1.0

	if errSpec != nil {
		return colorful.Color{}, errors.New("phong.Specular Failed: " + errSpec.Error())
	}

	if errDiff != nil {
		return colorful.Color{}, errors.New("phong.Diffuse Failed: " + errDiff.Error())
	}

	ambientContributionColor := utils.ScaleColor(ambientColor, ambientComponent)
	specularContributionColor := utils.ScaleColor(specularColor, specularComponent)
	diffuseContributionColor := utils.ScaleColor(diffuseColor, diffuseComponent)

	return CombineColors(CombineColors(ambientContributionColor, specularContributionColor), diffuseContributionColor), nil
}

package phong

import (
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/utils"
	"errors"
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/bcokert/render-cloud/model"
)

type PhongIlluminator struct {}

func Specular(reflectedVector, viewerVector mgl64.Vec3, shininess float64) (float64, error) {
	if !utils.IsDirectionVector(reflectedVector) {
		return 0, errors.New("phong.Specular requires reflectedVector to be a direction vector, received the vector [0, 0, 0]")
	}
	if !utils.IsDirectionVector(viewerVector) {
		return 0, errors.New("phong.Specular requires viewerVector to be a direction vector, received the vector [0, 0, 0]")
	}
	dot := reflectedVector.Normalize().Dot(viewerVector.Normalize())
	if dot < 0 {
		return 0, nil
	}
	return math.Pow(dot, shininess), nil
}

func Diffuse(specularComponent float64, lightVector, normalVector mgl64.Vec3) (float64, error) {
	if !utils.IsDirectionVector(lightVector) {
		return 0, errors.New("phong.Diffuse requires lightVector to be a direction vector, received the vector [0, 0, 0]")
	}
	if !utils.IsDirectionVector(normalVector) {
		return 0, errors.New("phong.Diffuse requires normalVector to be a direction vector, received the vector [0, 0, 0]")
	}
	return math.Max(0, (1 - specularComponent) * lightVector.Normalize().Dot(normalVector.Normalize())), nil
}

func ViewerVector(ray mgl64.Vec3) (mgl64.Vec3, error) {
	if !utils.IsDirectionVector(ray) {
		return mgl64.Vec3{}, errors.New("phong.ViewerVector requires ray to be a direction vector, received the vector [0, 0, 0]")
	}
	return ray.Mul(-1).Normalize(), nil
}

func LightVector(light model.Light) (mgl64.Vec3, error) {
	if !utils.IsDirectionVector(*light.Direction) {
		return mgl64.Vec3{}, errors.New("phong.LightVector requires light.Direction to be a direction vector, received the vector [0, 0, 0]")
	}
	return (*light.Direction).Mul(-1).Normalize(), nil
}

func ReflectedVector(normalVector, lightVector mgl64.Vec3) (mgl64.Vec3, error) {
	if !utils.IsDirectionVector(lightVector) {
		return mgl64.Vec3{}, errors.New("phong.LightVector requires lightVector to be a direction vector, received the vector [0, 0, 0]")
	}
	if !utils.IsDirectionVector(normalVector) {
		return mgl64.Vec3{}, errors.New("phong.LightVector requires normalVector to be a direction vector, received the vector [0, 0, 0]")
	}
	ln := lightVector.Normalize() //TODO: seems optional. Investigate
	nn := normalVector.Normalize()
	return nn.Mul((2*ln.Dot(nn))).Sub(ln).Normalize(), nil
}

func (this PhongIlluminator) CombineColors(c1, c2 colorful.Color) colorful.Color {
	return colorful.Color{
		c1.R + c2.R,
		c1.G + c2.G,
		c1.B + c2.B,
	}.Clamped()
}

func (this PhongIlluminator) MultiplyColors(c1, c2 colorful.Color) colorful.Color {
	return colorful.Color{
		c1.R * c2.R,
		c1.G * c2.G,
		c1.B * c2.B,
	}.Clamped()
}

//func (this PhongIlluminator)  IlluminateLocal(ambientColor, specularColor, diffuseColor colorful.Color, lightVector, normalVector, viewerVector, reflectedVector mgl64.Vec3, shininess float64) (colorful.Color, error) {
func (this PhongIlluminator)  IlluminateLocal(ray, normalVector mgl64.Vec3, material materials.Material, world model.World) (colorful.Color, error) {
	var err error

	resultColor := this.MultiplyColors(*world.Ambient, *material.Color)
	if world.Lights == nil {
		return resultColor, nil
	}

	var specularComponent, diffuseComponent float64
	var viewerVector, lightVector, reflectedVector mgl64.Vec3
	for _, light := range *world.Lights {

		if viewerVector, err = ViewerVector(ray); err != nil {
			return colorful.Color{}, errors.New("phong.ViewerVector Failed: " + err.Error())
		}

		if lightVector, err = LightVector(light); err != nil {
			return colorful.Color{}, errors.New("phong.LightVector Failed: " + err.Error())
		}

		if reflectedVector, err = ReflectedVector(normalVector, lightVector); err != nil {
			return colorful.Color{}, errors.New("phong.ReflectedVector Failed: " + err.Error())
		}

		if specularComponent, err = Specular(reflectedVector, viewerVector, *material.Shininess); err != nil {
			return colorful.Color{}, errors.New("phong.Specular Failed: " + err.Error())
		}

		if diffuseComponent, err = Diffuse(specularComponent, lightVector, normalVector); err != nil {
			return colorful.Color{}, errors.New("phong.Diffuse Failed: " + err.Error())
		}

		specularContributionColor := utils.ScaleColor(*light.Color, specularComponent)
		diffuseContributionColor := utils.ScaleColor(*material.Color, diffuseComponent)
		totalContribution := this.CombineColors(diffuseContributionColor, specularContributionColor)
		resultColor = this.CombineColors(resultColor, totalContribution)
	}

	return resultColor, nil
}

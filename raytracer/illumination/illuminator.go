package illumination

import (
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/bcokert/render-cloud/model"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
)

type Illuminator interface {
	IlluminateLocal(ray, normalVector mgl64.Vec3, material materials.Material, world model.World) (colorful.Color, error)
	CombineColors(c1, c2 colorful.Color) colorful.Color
	MultiplyColors(c1, c2 colorful.Color) colorful.Color
}

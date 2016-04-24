package primitives

import (
	"fmt"
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/bcokert/render-cloud/utils"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/tonestuff/quadratic"
	"math"
	"github.com/bcokert/render-cloud/validation"
)

type SpherePostRequest struct {
	Origin    *mgl64.Vec3                    `json:"origin" validate:"nonzero"`
	Material  materials.MaterialPostRequest `json:"material"`
	Radius    *float64                       `json:"radius"`
}

type SpheresPostRequest []SpherePostRequest

type Sphere struct {
	Origin    mgl64.Vec3          `json:"origin"`
	Material  materials.Material  `json:"material"`
	Radius    float64             `json:"radius" validate:"nonzero,min=0"`
}

func (this *Sphere) FromPostRequest(validator validation.Validator, postRequest SpherePostRequest) error {
	if err := validator.Validate(postRequest); err != nil {
		return err
	}

	this.Origin = utils.DefaultVector(postRequest.Origin, mgl64.Vec3{0, 0, 0})
	if err := this.Material.FromPostRequest(validator, postRequest.Material); err != nil {
		return fmt.Errorf("Material: not a valid Material (%s)", err.Error())
	}
	this.Radius = utils.DefaultFloat(postRequest.Radius, 1)

	return validator.Validate(this)
}

func (this Sphere) GetMaterial() materials.Material {
	return this.Material
}

func (this Sphere) GetNormalAtPoint(point mgl64.Vec3) (mgl64.Vec3, error) {
	normal := point.Sub(this.Origin)
	if mgl64.FloatEqual(normal.Len(), this.Radius) {
		return normal.Normalize(), nil
	} else {
		return mgl64.Vec3{}, fmt.Errorf("Cannot get normal at %v. Point must be on surface of sphere.", point)
	}
}

func (this Sphere) FindClosestRayCollision(origin mgl64.Vec3, direction mgl64.Vec3) *float64 {
	originMinusSphere := origin.Sub(this.Origin)
	direction = direction.Normalize()

	a := complex(1.0, 0)
	b := complex(2*originMinusSphere.Dot(direction), 0)
	c := complex(originMinusSphere.Dot(originMinusSphere)-this.Radius*this.Radius, 0)

	ans1, ans2 := quadratic.Solve(a, b, c)

	if imag(ans1) != 0 || imag(ans2) != 0 {
		return nil
	}

	distance := math.Min(real(ans1), real(ans2))
	return &distance
}

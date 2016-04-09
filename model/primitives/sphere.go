package primitives

import (
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/tonestuff/quadratic"
	"math"
	"errors"
	"fmt"
)

type Sphere struct {
	Origin    *mgl64.Vec3 `json:"origin,omitempty"`
	Radius    *float64        `json:"radius,omitempty"`
	Material  *materials.Material `json:"material,omitempty"`
}

func (sphere Sphere) GetOrigin() mgl64.Vec3 {
	return *sphere.Origin
}

func (sphere Sphere) GetRadius() float64 {
	return *sphere.Radius
}

func (sphere Sphere) GetMaterial() materials.Material {
	return *sphere.Material
}

func (sphere Sphere) GetNormalAtPoint(point mgl64.Vec3) (mgl64.Vec3, error) {
	normal := point.Sub(sphere.GetOrigin())
	if mgl64.FloatEqual(normal.Len(), sphere.GetRadius()) {
		return normal.Normalize(), nil
	} else {
		return mgl64.Vec3{}, errors.New(fmt.Sprintf("Cannot get normal at %v. Point must be on surface of sphere.", point))
	}
}

func (sphere Sphere) FindClosestRayCollision(origin mgl64.Vec3, direction mgl64.Vec3) *float64 {
	originMinusSphere := origin.Sub(sphere.GetOrigin())
	direction = direction.Normalize()

	a := complex(1.0, 0)
	b := complex(2 * originMinusSphere.Dot(direction), 0)
	c := complex(originMinusSphere.Dot(originMinusSphere) - sphere.GetRadius() * sphere.GetRadius(), 0)

	ans1, ans2 := quadratic.Solve(a, b, c)

	if imag(ans1) != 0 || imag(ans2) != 0 {
		return nil
	}

	distance := math.Min(real(ans1), real(ans2))
	return &distance
}

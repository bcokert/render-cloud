package primitives

import (
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/go-gl/mathgl/mgl64"
)

type Primitive interface {
	FindClosestRayCollision(mgl64.Vec3, mgl64.Vec3) mgl64.Vec3
	GetMaterial() materials.Material
	GetOrigin() mgl64.Vec3
}

package primitives

import (
	"github.com/bcokert/render-cloud/vector"
	"github.com/bcokert/render-cloud/model/materials"
)

type Primitive interface {
	FindClosestRayCollision(vector.Vector3, vector.Vector3) vector.Vector3
	GetMaterial() materials.Material
	GetOrigin() vector.Vector3
}

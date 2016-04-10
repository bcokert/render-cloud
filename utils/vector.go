package utils

import (
	"github.com/go-gl/mathgl/mgl64"
)

func IsDirectionVector(vector mgl64.Vec3) bool {
	return vector.X() != 0 || vector.Y() != 0 || vector.Z() != 0
}

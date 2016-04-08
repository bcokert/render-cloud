package primitives_test

import (
	"github.com/bcokert/render-cloud/testutils"
	"github.com/bcokert/render-cloud/utils"
	"github.com/lucasb-eyer/go-colorful"
	"testing"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/go-gl/mathgl/mgl64"
)

func TestSphereJsonEncodes(t *testing.T) {
	sphere := primitives.Sphere{Origin: &mgl64.Vec3{0, 0, 0}, Radius: utils.FloatPointer(5.0), Material: &materials.Material{Color: &colorful.Color{R: 1, G: 0, B: 0}, Shininess: utils.FloatPointer(9)}}

	expectedJson := "{\"origin\":[0,0,0],\"radius\":5,\"material\":{\"color\":{\"R\":1,\"G\":0,\"B\":0},\"shininess\":9}}"

	testutils.ExpectJsonEncoding(t, &sphere, expectedJson)
}

func TestFindClosestRayCollision(t *testing.T) {
	cases := []struct{
		Sphere primitives.Sphere
		Origin mgl64.Vec3
		Direction mgl64.Vec3
		ExpectedResult *mgl64.Vec3
	}{
		{primitives.Sphere{Origin: &mgl64.Vec3{0, 0, 0}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, &mgl64.Vec3{0, 0, -1}},
		{primitives.Sphere{Origin: &mgl64.Vec3{0, 0, 0}, Radius: utils.FloatPointer(3.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, &mgl64.Vec3{0, 0, -3}},
		{primitives.Sphere{Origin: &mgl64.Vec3{0, 1, 0}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, &mgl64.Vec3{0, 0, 0}},
		{primitives.Sphere{Origin: &mgl64.Vec3{1, 0, 0}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, &mgl64.Vec3{0, 0, 0}},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{1,1,1}, &mgl64.Vec3{4.422649730810371, 4.422649730810371, 4.422649730810371}},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{1,0,0}, nil},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{0,1,0}, nil},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{0,0,1}, nil},
	}

	for i, testCase := range cases {
		res := testCase.Sphere.FindClosestRayCollision(testCase.Origin, testCase.Direction)
		if (testCase.ExpectedResult == nil && res != nil) || testCase.ExpectedResult != nil && !(mgl64.FloatEqual(res.X(), testCase.ExpectedResult.X()) && mgl64.FloatEqual(res.Y(), testCase.ExpectedResult.Y()) && mgl64.FloatEqual(res.Z(), testCase.ExpectedResult.Z())) {
			t.Errorf("Failed to find correct collision for testCase %d. Expected %s, found %s", i, testCase.ExpectedResult, res)
		}
	}
}

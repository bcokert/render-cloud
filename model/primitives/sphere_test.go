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

func TestSphereIsAPrimitive(t *testing.T) {
	primArray := []primitives.Primitive{}
	primArray = append(primArray, primitives.Sphere{}) // type error if not a primitive
}

func TestGetNormalAtPoint(t *testing.T) {
	testCases := []struct{
	    Sphere primitives.Sphere
		Point mgl64.Vec3
		Expected mgl64.Vec3
	}{
	    {primitives.Sphere{Origin: &mgl64.Vec3{0,0,0}, Radius: utils.FloatPointer(1), Material: nil}, mgl64.Vec3{1,0,0}, mgl64.Vec3{1,0,0}},
	    {primitives.Sphere{Origin: &mgl64.Vec3{0,0,0}, Radius: utils.FloatPointer(2), Material: nil}, mgl64.Vec3{0,2,0}, mgl64.Vec3{0,1,0}},
	    {primitives.Sphere{Origin: &mgl64.Vec3{0,0,0}, Radius: utils.FloatPointer(0.5), Material: nil}, mgl64.Vec3{0,0,0.5}, mgl64.Vec3{0,0,1}},

	    {primitives.Sphere{Origin: &mgl64.Vec3{3,3,3}, Radius: utils.FloatPointer(2.5), Material: nil}, mgl64.Vec3{3,3,5.5}, mgl64.Vec3{0,0,1}},
	    {primitives.Sphere{Origin: &mgl64.Vec3{3,3,3}, Radius: utils.FloatPointer(2.5), Material: nil}, mgl64.Vec3{3,5.5,3}, mgl64.Vec3{0,1,0}},
	    {primitives.Sphere{Origin: &mgl64.Vec3{9,3,1}, Radius: utils.FloatPointer(4.25), Material: nil}, mgl64.Vec3{13.25,3,1}, mgl64.Vec3{1,0,0}},

	    {primitives.Sphere{Origin: &mgl64.Vec3{4,2,-5}, Radius: utils.FloatPointer(2.11), Material: nil}, mgl64.Vec3{4,2,-5}.Add(mgl64.Vec3{5,2,1}.Normalize().Mul(2.11)), mgl64.Vec3{5,2,1}.Normalize()},
	    {primitives.Sphere{Origin: &mgl64.Vec3{-3,1.22,5}, Radius: utils.FloatPointer(4.05), Material: nil}, mgl64.Vec3{-3,1.22,5}.Add(mgl64.Vec3{2,1,1}.Normalize().Mul(4.05)), mgl64.Vec3{2,1,1}.Normalize()},
	}

	for i, testCase := range testCases {
		normal, err := testCase.Sphere.GetNormalAtPoint(testCase.Point)
		if err != nil {
			t.Errorf("Case %d failed. Threw unexpected error %s", i, err.Error())
		}

		if !normal.ApproxEqual(testCase.Expected) {
			t.Errorf("Case %d failed. Expected %v, received %v", i, testCase.Expected, normal)
		}
	}
}

func TestGetNormalAtPointError(t *testing.T) {
	testCases := []struct{
		Sphere primitives.Sphere
		Point mgl64.Vec3
		Expected string
	}{
		{primitives.Sphere{Origin: &mgl64.Vec3{0,0,0}, Radius: utils.FloatPointer(1), Material: nil}, mgl64.Vec3{0,0,0}, "Cannot get normal at [0 0 0]. Point must be on surface of sphere."},
		{primitives.Sphere{Origin: &mgl64.Vec3{0,0,0}, Radius: utils.FloatPointer(2), Material: nil}, mgl64.Vec3{0,2.1,0}, "Cannot get normal at [0 2.1 0]. Point must be on surface of sphere."},

		{primitives.Sphere{Origin: &mgl64.Vec3{4,2,-5}, Radius: utils.FloatPointer(2.11), Material: nil}, mgl64.Vec3{4,2,-5}.Add(mgl64.Vec3{5,2,1}), "Cannot get normal at [9 4 -4]. Point must be on surface of sphere."},
		{primitives.Sphere{Origin: &mgl64.Vec3{-3,1.22,5}, Radius: utils.FloatPointer(4.05), Material: nil}, mgl64.Vec3{-3,1.23,5}.Add(mgl64.Vec3{2,1,1}), "Cannot get normal at [-1 2.23 6]. Point must be on surface of sphere."},
	}

	for i, testCase := range testCases {
		_, err := testCase.Sphere.GetNormalAtPoint(testCase.Point)
		if err == nil {
			t.Errorf("Case %d failed. Expected an error but received none", i)
		}

		if err.Error() != testCase.Expected {
			t.Errorf("Case %d failed. Expected error %s, received %s", i, testCase.Expected, err.Error())
		}
	}
}

func TestFindClosestRayCollision(t *testing.T) {
	cases := []struct{
		Sphere primitives.Sphere
		Origin mgl64.Vec3
		Direction mgl64.Vec3
		ExpectedResult *float64
	}{
		{primitives.Sphere{Origin: &mgl64.Vec3{0, 0, 0}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, utils.FloatPointer(9)},
		{primitives.Sphere{Origin: &mgl64.Vec3{0, 0, 0}, Radius: utils.FloatPointer(3.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, utils.FloatPointer(7)},
		{primitives.Sphere{Origin: &mgl64.Vec3{0, 1, 0}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, utils.FloatPointer(10)},
		{primitives.Sphere{Origin: &mgl64.Vec3{1, 0, 0}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, -10}, mgl64.Vec3{0,0,1}, utils.FloatPointer(10)},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{1,1,1}, utils.FloatPointer(7.66025403784438)},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{1,0,0}, nil},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{0,1,0}, nil},
		{primitives.Sphere{Origin: &mgl64.Vec3{5, 5, 5}, Radius: utils.FloatPointer(1.0), Material: nil}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{0,0,1}, nil},
	}

	for i, testCase := range cases {
		res := testCase.Sphere.FindClosestRayCollision(testCase.Origin, testCase.Direction)
		if (testCase.ExpectedResult == nil && res != nil) || testCase.ExpectedResult != nil && *res != *testCase.ExpectedResult {
			t.Errorf("Failed to find correct collision for testCase %d. Expected %s, found %s", i, *testCase.ExpectedResult, *res)
		}
	}
}

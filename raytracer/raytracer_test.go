package raytracer_test

import (
	"testing"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/bcokert/render-cloud/utils"
	"github.com/bcokert/render-cloud/raytracer"
	"github.com/bcokert/render-cloud/model/materials"
	"github.com/bcokert/render-cloud/model"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/raytracer/illumination"
	"github.com/bcokert/render-cloud/raytracer/illumination/phong"
	"errors"
)

func TestGetClosestCollisionAndPrimitive(t *testing.T) {
    testCases := map[string]struct{
	    Origin mgl64.Vec3
	    Ray mgl64.Vec3
	    Objects []primitives.Primitive

        ExpectedPoint mgl64.Vec3
	    ExpectedPrimitiveIndex *int
    }{
        "Sphere directly in front": {
	        Origin: mgl64.Vec3{0,0,0},
	        Ray: mgl64.Vec3{0,0,1},
	        Objects: []primitives.Primitive{
		        primitives.Sphere{
			        Origin: &mgl64.Vec3{0,0,10},
			        Radius: utils.FloatPointer(1),
			        Material: &materials.Material{},
		        },
	        },
	        ExpectedPoint: mgl64.Vec3{0,0,9},
	        ExpectedPrimitiveIndex: utils.IntPointer(0),
        },
	    "3 Spheres in a row": {
		    Origin: mgl64.Vec3{0,0,0},
		    Ray: mgl64.Vec3{0,0,1},
		    Objects: []primitives.Primitive{
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{0,0,10},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{0,0,6},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{0,0,15},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
		    },
		    ExpectedPoint: mgl64.Vec3{0,0,5},
		    ExpectedPrimitiveIndex: utils.IntPointer(1),
	    },
	    "3 Spheres all missed": {
		    Origin: mgl64.Vec3{0,0,0},
		    Ray: mgl64.Vec3{0,0,1},
		    Objects: []primitives.Primitive{
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{3,2,10},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{-6,-1,6},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{0,4,15},
				    Radius: utils.FloatPointer(2),
				    Material: &materials.Material{},
			    },
		    },
		    ExpectedPoint: mgl64.Vec3{},
		    ExpectedPrimitiveIndex: nil,
	    },
	    "3 Spheres 1 hit": {
		    Origin: mgl64.Vec3{0,0,0},
		    Ray: mgl64.Vec3{0,0,1},
		    Objects: []primitives.Primitive{
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{3,2,10},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{-6,-1,6},
				    Radius: utils.FloatPointer(1),
				    Material: &materials.Material{},
			    },
			    primitives.Sphere{
				    Origin: &mgl64.Vec3{0,0,15},
				    Radius: utils.FloatPointer(2),
				    Material: &materials.Material{},
			    },
		    },
		    ExpectedPoint: mgl64.Vec3{0,0,13},
		    ExpectedPrimitiveIndex: utils.IntPointer(2),
	    },
	    "no Spheres": {
		    Origin: mgl64.Vec3{0,0,0},
		    Ray: mgl64.Vec3{0,0,1},
		    Objects: []primitives.Primitive{},
		    ExpectedPoint: mgl64.Vec3{},
		    ExpectedPrimitiveIndex: nil,
	    },
    }

    for name, testCase := range testCases {
        collision, primitive := raytracer.DefaultRaytracer{}.GetClosestCollisionAndPrimitive(testCase.Origin, testCase.Ray, testCase.Objects)
        if !testCase.ExpectedPoint.ApproxEqual(collision) {
            t.Errorf("'%s' failed. Expected collision %v, received %v", name, testCase.ExpectedPoint, collision)
        }
        if testCase.ExpectedPrimitiveIndex == nil && primitive != nil {
            t.Errorf("'%s' failed. Expected primtive to be nil, received %d", name, (*primitive).String())
        }
	    if testCase.ExpectedPrimitiveIndex != nil && primitive == nil {
		    t.Errorf("'%s' failed. Expected primtive to be %s, received nil", name, testCase.Objects[*testCase.ExpectedPrimitiveIndex].String())
	    } else if (testCase.ExpectedPrimitiveIndex != nil && (*primitive).String() != testCase.Objects[*testCase.ExpectedPrimitiveIndex].String()) {
		    t.Errorf("'%s' failed. Expected primtive to be %s, received %s", name, testCase.Objects[*testCase.ExpectedPrimitiveIndex].String(), (*primitive).String())
	    }
    }
}

func TestGetScreenVectors(t *testing.T) {
    testCases := map[string]struct{
	    Camera model.Camera
        ExpectedUp mgl64.Vec3
        ExpectedRight mgl64.Vec3
        ExpectedTopLeft mgl64.Vec3
    }{
        "Basic 1": {
	        Camera: model.Camera{
		        Origin:         &mgl64.Vec3{0,0,-10},
		        Direction:      &mgl64.Vec3{0,0,1},
		        Up:             &mgl64.Vec3{0,1,0},
		        ScreenWidth:    utils.FloatPointer(2),
		        ScreenHeight:   utils.FloatPointer(2),
		        ScreenDistance: utils.FloatPointer(1),
	        },
	        ExpectedUp: mgl64.Vec3{0,1,0},
	        ExpectedRight: mgl64.Vec3{1,0,0},
	        ExpectedTopLeft: mgl64.Vec3{-1,1,-9},
        },
	    "Scaled": {
		    Camera: model.Camera{
			    Origin:         &mgl64.Vec3{0,0,-9},
			    Direction:      &mgl64.Vec3{0,0,4},
			    Up:             &mgl64.Vec3{0,2,0},
			    ScreenWidth:    utils.FloatPointer(2),
			    ScreenHeight:   utils.FloatPointer(2),
			    ScreenDistance: utils.FloatPointer(2),
		    },
		    ExpectedUp: mgl64.Vec3{0,1,0},
		    ExpectedRight: mgl64.Vec3{1,0,0},
		    ExpectedTopLeft: mgl64.Vec3{-1,1,-7},
	    },
	    "Rotated": {
		    Camera: model.Camera{
			    Origin:         &mgl64.Vec3{0,0,-10},
			    Direction:      &mgl64.Vec3{0,0,-1},
			    Up:             &mgl64.Vec3{0,-1,0},
			    ScreenWidth:    utils.FloatPointer(2),
			    ScreenHeight:   utils.FloatPointer(2),
			    ScreenDistance: utils.FloatPointer(1),
		    },
		    ExpectedUp: mgl64.Vec3{0,-1,0},
		    ExpectedRight: mgl64.Vec3{1,0,0},
		    ExpectedTopLeft: mgl64.Vec3{-1,-1,-11},
	    },
    }

    for name, testCase := range testCases {
        up, right, topLeft := raytracer.DefaultRaytracer{}.GetScreenVectors(testCase.Camera)
        if !up.ApproxEqual(testCase.ExpectedUp) {
            t.Errorf("'%s' failed. Expected up to be %v, received %v", name, testCase.ExpectedUp, up)
        }
	    if !right.ApproxEqual(testCase.ExpectedRight) {
		    t.Errorf("'%s' failed. Expected right to be %v, received %v", name, testCase.ExpectedRight, right)
	    }
	    if !topLeft.ApproxEqual(testCase.ExpectedTopLeft) {
		    t.Errorf("'%s' failed. Expected topLeft to be %v, received %v", name, testCase.ExpectedTopLeft, topLeft)
	    }
    }
}

func Test(t *testing.T) {
    testCases := map[string]struct{
	    XYPairs []float64
	    Width, Height float64
	    Camera model.Camera
	    Up, Right, TopLeft mgl64.Vec3
        ExpectedRays []mgl64.Vec3
    }{
        "Basic in front": {
	        XYPairs: []float64{0,0, 1,0, 0,1, 1,1},
	        Width: 2,
	        Height: 2,
	        Camera: model.Camera{
		        Origin:         &mgl64.Vec3{0,0,-10},
		        Direction:      &mgl64.Vec3{0,0,1},
		        Up:             &mgl64.Vec3{0,1,0},
		        ScreenWidth:    utils.FloatPointer(2),
		        ScreenHeight:   utils.FloatPointer(2),
		        ScreenDistance: utils.FloatPointer(1),
	        },
	        ExpectedRays: []mgl64.Vec3{
		        mgl64.Vec3{-0.5,0.5,1}.Normalize(),
		        mgl64.Vec3{0.5,0.5,1}.Normalize(),
		        mgl64.Vec3{-0.5,-0.5,1}.Normalize(),
		        mgl64.Vec3{0.5,-0.5,1}.Normalize(),
	        },
        },
	    "Rotated": {
		    XYPairs: []float64{0,0, 1,0, 0,1, 1,1},
		    Width: 2,
		    Height: 2,
		    Camera: model.Camera{
			    Origin:         &mgl64.Vec3{0,0,-10},
			    Direction:      &mgl64.Vec3{0,0,1},
			    Up:             &mgl64.Vec3{0,-1,0},
			    ScreenWidth:    utils.FloatPointer(2),
			    ScreenHeight:   utils.FloatPointer(2),
			    ScreenDistance: utils.FloatPointer(1),
		    },
		    ExpectedRays: []mgl64.Vec3{
			    mgl64.Vec3{0.5,-0.5,1}.Normalize(),
			    mgl64.Vec3{-0.5,-0.5,1}.Normalize(),
			    mgl64.Vec3{0.5,0.5,1}.Normalize(),
			    mgl64.Vec3{-0.5,0.5,1}.Normalize(),
		    },
	    },
	    "Large, far screen, 4:3 ratio, non-normal directions": {
		    XYPairs: []float64{0,0, 1,0, 0,1, 1,1},
		    Width: 2,
		    Height: 2,
		    Camera: model.Camera{
			    Origin:         &mgl64.Vec3{0,0,-10},
			    Direction:      &mgl64.Vec3{0,0,51},
			    Up:             &mgl64.Vec3{0,4,0},
			    ScreenWidth:    utils.FloatPointer(4),
			    ScreenHeight:   utils.FloatPointer(3),
			    ScreenDistance: utils.FloatPointer(5),
		    },
		    ExpectedRays: []mgl64.Vec3{
			    mgl64.Vec3{-1,0.75,5}.Normalize(),
			    mgl64.Vec3{1,0.75,5}.Normalize(),
			    mgl64.Vec3{-1,-0.75,5}.Normalize(),
			    mgl64.Vec3{1,-0.75,5}.Normalize(),
		    },
	    },
    }

	testRaytracer := raytracer.DefaultRaytracer{}
    for name, testCase := range testCases {
	    up, right, topLeft := testRaytracer.GetScreenVectors(testCase.Camera)
	    for i := 0; i<len(testCase.ExpectedRays); i++ {
		    ray := testRaytracer.GetRayForPixel(testCase.XYPairs[i*2], testCase.XYPairs[i*2+1], testCase.Width, testCase.Height, testCase.Camera, up, right, topLeft)
		    if !ray.ApproxEqual(testCase.ExpectedRays[i]) {
			    t.Errorf("'%s' failed for %v,%v. Expected %v, received %v", name, testCase.XYPairs[i*2], testCase.XYPairs[i*2+1], testCase.ExpectedRays[i], ray)
		    }
	    }
    }
}

type illuminatorAlwaysRed struct {
	phong.PhongIlluminator
}

func (this illuminatorAlwaysRed) IlluminateLocal(ray, normalVector mgl64.Vec3, material materials.Material, world model.World) (colorful.Color, error) {
	return colorful.Color{1,0,0}, nil
}

type illuminatorAlwaysFails struct {
	phong.PhongIlluminator
}

func (this illuminatorAlwaysFails) IlluminateLocal(ray, normalVector mgl64.Vec3, material materials.Material, world model.World) (colorful.Color, error) {
	return colorful.Color{}, errors.New("I failed to illuminate you")
}

func TestTraceScene(t *testing.T) {
    testCases := map[string]struct{
	    Scene model.Scene
	    Illuminator illumination.Illuminator
	    Width, Height uint
        Expected []colorful.Color
    }{
        "No Lights": {
	        Scene: model.Scene{
		        utils.UintPointer(2345),
		        &model.World{
			        Ambient: &colorful.Color{0.2,0.2,0.2},
			        Background: &colorful.Color{0.1,0.1,0.1},
			        Camera: &model.Camera{
				        Origin: &mgl64.Vec3{0,0,-8},
				        Direction: &mgl64.Vec3{0,0,1},
				        Up: &mgl64.Vec3{0,1,0},
				        ScreenWidth: utils.FloatPointer(5),
				        ScreenHeight: utils.FloatPointer(5),
				        ScreenDistance: utils.FloatPointer(2),
			        },
			        Lights: &[]model.Light{
			        },
		        },
		        &[]primitives.Sphere{
			        primitives.Sphere{
				        Origin: &mgl64.Vec3{0,0,0},
				        Radius: utils.FloatPointer(4),
				        Material: &materials.Material{
					        Color: &colorful.Color{0.7,0.1,0.1},
					        Shininess: utils.FloatPointer(10),
				        },
			        },
		        },
	        },
	        Illuminator: phong.PhongIlluminator{},
	        Width: 4,
	        Height: 4,
	        Expected: []colorful.Color{
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},//2
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},//7
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},//12
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
		        colorful.Color{0.1,0.1,0.1},
	        },
        },
	    "Collisions always red": {
		    Scene: model.Scene{
			    utils.UintPointer(2345),
			    &model.World{
				    Ambient: &colorful.Color{0.2,0.2,0.2},
				    Background: &colorful.Color{0.1,0.1,0.1},
				    Camera: &model.Camera{
					    Origin: &mgl64.Vec3{0,0,-8},
					    Direction: &mgl64.Vec3{0,0,1},
					    Up: &mgl64.Vec3{0,1,0},
					    ScreenWidth: utils.FloatPointer(5),
					    ScreenHeight: utils.FloatPointer(5),
					    ScreenDistance: utils.FloatPointer(2),
				    },
				    Lights: &[]model.Light{
					    model.Light{
						    Direction: &mgl64.Vec3{1, 1, 1},
						    Color: &colorful.Color{0.4, 0.4, 0.4},
					    },
				    },
			    },
			    &[]primitives.Sphere{
				    primitives.Sphere{
					    Origin: &mgl64.Vec3{0,0,0},
					    Radius: utils.FloatPointer(4),
					    Material: &materials.Material{
						    Color: &colorful.Color{0.7,0.1,0.1},
						    Shininess: utils.FloatPointer(10),
					    },
				    },
			    },
		    },
		    Illuminator: illuminatorAlwaysRed{},
		    Width: 4,
		    Height: 4,
		    Expected: []colorful.Color{
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//2
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{1,0,0},
			    colorful.Color{1,0,0},
			    colorful.Color{0.1,0.1,0.1},//7
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{1,0,0},
			    colorful.Color{1,0,0},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//12
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
		    },
	    },
	    "No Collisions": {
		    Scene: model.Scene{
			    utils.UintPointer(2345),
			    &model.World{
				    Ambient: &colorful.Color{0.2,0.2,0.2},
				    Background: &colorful.Color{0.1,0.1,0.1},
				    Camera: &model.Camera{
					    Origin: &mgl64.Vec3{0,0,-8},
					    Direction: &mgl64.Vec3{0,0,1},
					    Up: &mgl64.Vec3{0,1,0},
					    ScreenWidth: utils.FloatPointer(5),
					    ScreenHeight: utils.FloatPointer(5),
					    ScreenDistance: utils.FloatPointer(2),
				    },
				    Lights: &[]model.Light{
					    model.Light{
						    Direction: &mgl64.Vec3{1, -1, 1},
						    Color: &colorful.Color{0.4, 0.4, 0.4},
					    },
				    },
			    },
			    &[]primitives.Sphere{
				    primitives.Sphere{
					    Origin: &mgl64.Vec3{0,100,0},
					    Radius: utils.FloatPointer(4),
					    Material: &materials.Material{
						    Color: &colorful.Color{0.7,0.1,0.1},
						    Shininess: utils.FloatPointer(10),
					    },
				    },
			    },
		    },
		    Illuminator: phong.PhongIlluminator{},
		    Width: 4,
		    Height: 4,
		    Expected: []colorful.Color{
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//2
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//7
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//12
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
		    },
	    },
	    "Default": {
		    Scene: model.Scene{
			    utils.UintPointer(2345),
			    &model.World{
				    Ambient: &colorful.Color{0.2,0.2,0.2},
				    Background: &colorful.Color{0.1,0.1,0.1},
				    Camera: &model.Camera{
					    Origin: &mgl64.Vec3{0,0,-8},
					    Direction: &mgl64.Vec3{0,0,1},
					    Up: &mgl64.Vec3{0,1,0},
					    ScreenWidth: utils.FloatPointer(5),
					    ScreenHeight: utils.FloatPointer(5),
					    ScreenDistance: utils.FloatPointer(2),
				    },
				    Lights: &[]model.Light{
					    model.Light{
						    Direction: &mgl64.Vec3{1, -1, 1},
						    Color: &colorful.Color{0.4, 0.4, 0.4},
					    },
				    },
			    },
			    &[]primitives.Sphere{
				    primitives.Sphere{
					    Origin: &mgl64.Vec3{0,0,0},
					    Radius: utils.FloatPointer(4),
					    Material: &materials.Material{
						    Color: &colorful.Color{0.7,0.1,0.1},
						    Shininess: utils.FloatPointer(10),
					    },
				    },
			    },
		    },
		    Illuminator: phong.PhongIlluminator{},
		    Width: 4,
		    Height: 4,
		    Expected: []colorful.Color{
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//2
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.7162520817103684, 0.18944325611395135, 0.18944325611395135},
			    colorful.Color{0.489638526068923, 0.06994836086719207, 0.06994836086719207},
			    colorful.Color{0.1,0.1,0.1},//7
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.489638526068923, 0.06994836086719207, 0.06994836086719207},
			    colorful.Color{0.20298111932102886, 0.028997302760146983, 0.028997302760146983},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},//12
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
			    colorful.Color{0.1,0.1,0.1},
		    },
	    },
    }

    errorTestCases := map[string]struct{
	    Scene model.Scene
	    Illuminator illumination.Illuminator
	    Width, Height uint
        Expected string
    }{
        "Illuminating Fails": {
	        Scene: model.Scene{
		        utils.UintPointer(2345),
		        &model.World{
			        Ambient: &colorful.Color{0.2,0.2,0.2},
			        Background: &colorful.Color{0.1,0.1,0.1},
			        Camera: &model.Camera{
				        Origin: &mgl64.Vec3{0,0,-8},
				        Direction: &mgl64.Vec3{0,0,1},
				        Up: &mgl64.Vec3{0,1,0},
				        ScreenWidth: utils.FloatPointer(5),
				        ScreenHeight: utils.FloatPointer(5),
				        ScreenDistance: utils.FloatPointer(2),
			        },
			        Lights: &[]model.Light{
				        model.Light{
					        Direction: &mgl64.Vec3{1, -1, 1},
					        Color: &colorful.Color{0.4, 0.4, 0.4},
				        },
			        },
		        },
		        &[]primitives.Sphere{
			        primitives.Sphere{
				        Origin: &mgl64.Vec3{0,0,0},
				        Radius: utils.FloatPointer(4),
				        Material: &materials.Material{
					        Color: &colorful.Color{0.7,0.1,0.1},
					        Shininess: utils.FloatPointer(10),
				        },
			        },
		        },
	        },
	        Illuminator: illuminatorAlwaysFails{},
	        Width: 4,
	        Height: 4,
	        Expected: "An error occurred tracing pixel x:1, y:1, while finding local illumination. Error: I failed to illuminate you",
        },
    }

    for name, testCase := range testCases {
        result, err := raytracer.DefaultRaytracer{}.TraceScene(testCase.Scene, testCase.Illuminator, testCase.Width, testCase.Height)
        if err != nil {
            t.Errorf("'%s' failed. Unexpected error %s", name, err.Error())
        }
	    for i, pixel := range result {
		    if !pixel.AlmostEqualRgb(testCase.Expected[i]) {
			    t.Errorf("'%s' failed for pixel %d. Expected %v, received %v", name, i, testCase.Expected[i], pixel)
		    }
	    }
    }

    for name, testCase := range errorTestCases {
        _, err := raytracer.DefaultRaytracer{}.TraceScene(testCase.Scene, testCase.Illuminator, testCase.Width, testCase.Height)
        if err == nil {
            t.Errorf("'%s' failed. Expected error but none occured", name)
        }
        if err.Error() != testCase.Expected {
            t.Errorf("'%s' failed. Expected %s, received %s", name, testCase.Expected, err.Error())
        }
    }
}

package raytracer

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/bcokert/render-cloud/raytracer/illumination"
	"errors"
	"fmt"
)

type Raytracer interface {
	TraceScene(scene model.Scene, illuminator illumination.Illuminator, width, height uint) ([]colorful.Color, error)
	GetClosestCollisionAndPrimitive(origin, ray mgl64.Vec3, objects []primitives.Primitive) (collisionPoint mgl64.Vec3, primitive *primitives.Primitive)
	GetScreenVectors(camera model.Camera) (up, right, topLeft mgl64.Vec3)
	GetRayForPixel(x, y, width, height float64, camera model.Camera, screenUp, screenRight, screenTopLeft mgl64.Vec3) (ray mgl64.Vec3)
}

type DefaultRaytracer struct {}

func (this DefaultRaytracer) GetClosestCollisionAndPrimitive(origin, ray mgl64.Vec3, objects []primitives.Primitive) (collisionPoint mgl64.Vec3, primitive *primitives.Primitive) {
	var distance float64
	for _, prim := range objects {
		collision := prim.FindClosestRayCollision(origin, ray)
		if collision != nil {
			if primitive == nil || *collision < distance {
				distance = *collision
				foundPrimitive := prim // copy prim, since it's reassigned in each iteration
				primitive = &foundPrimitive
			}
		}
	}
	return origin.Add(ray.Mul(distance)), primitive
}

func (this DefaultRaytracer) GetScreenVectors(camera model.Camera) (up, right, topLeft mgl64.Vec3) {
	forward := camera.GetDirection().Normalize()
	up = camera.GetUp().Normalize()
	right = up.Cross(forward).Normalize()
	center := (*camera.Origin).Add(forward.Mul(camera.GetScreenDistance()))
	topLeft = center.Add(up.Mul(camera.GetScreenHeight() / 2.0)).Sub(right.Mul(camera.GetScreenWidth() / 2.0))
	return up, right, topLeft
}

func (this DefaultRaytracer) GetRayForPixel(x, y, width, height float64, camera model.Camera, screenUp, screenRight, screenTopLeft mgl64.Vec3) (ray mgl64.Vec3) {
	widthOfPixel := camera.GetScreenWidth() / width
	halfWidthOfPixel := widthOfPixel/2
	heightOfPixel := camera.GetScreenHeight() / height
	halfHeightOfPixel := heightOfPixel/2
	dx := x * widthOfPixel + halfWidthOfPixel
	dy := y * heightOfPixel + halfHeightOfPixel
	screenPoint := screenTopLeft.Add(screenRight.Mul(dx)).Sub(screenUp.Mul(dy))
	ray = screenPoint.Sub(*camera.Origin).Normalize()
	return ray
}

func (this DefaultRaytracer) TraceScene(scene model.Scene, illuminator illumination.Illuminator, width, height uint) ([]colorful.Color, error) {
	screenUp, screenRight, screenTopLeft := this.GetScreenVectors(scene.GetWorld().GetCamera())

	var x, y, ww, hh float64 = 0, 0, float64(width), float64(height)
	var colors []colorful.Color
	for y = 0; y < hh; y+=1 {
		for x = 0; x < ww; x+=1 {

			// Optimization: no lights means collisions don't matter
			if len(scene.GetWorld().GetLights()) == 0 {
				colors = append(colors, scene.GetWorld().GetBackground())
				continue
			}

			// Create a ray through the screen at x,y and check for collisions
			ray := this.GetRayForPixel(x, y, ww, hh, scene.GetWorld().GetCamera(), screenUp, screenRight, screenTopLeft)
			primitiveObjects := make([]primitives.Primitive, len(scene.GetSpheres()), len(scene.GetSpheres()))
			for i, sphere := range scene.GetSpheres() {
				// go does not implicitly convert slices into an interface type, so we must do it manually
				primitiveObjects[i] = sphere
			}
			collisionPoint, primitive := this.GetClosestCollisionAndPrimitive(scene.GetWorld().GetCamera().GetOrigin(), ray, primitiveObjects)

			// If collisions, do local + global illumination. Else, return background color
			if primitive == nil {
				colors = append(colors, scene.GetWorld().GetBackground())
			} else {
				normal, err := (*primitive).GetNormalAtPoint(collisionPoint)
				if err != nil {
					return []colorful.Color{}, errors.New(fmt.Sprintf("An error occurred tracing pixel x:%v, y:%v, while finding surface normal. Error: %s", x, y, err.Error()))
				}
				resultColor, err := illuminator.IlluminateLocal(ray, normal, (*primitive).GetMaterial(), scene.GetWorld())
				if err != nil {
					return []colorful.Color{}, errors.New(fmt.Sprintf("An error occurred tracing pixel x:%v, y:%v, while finding local illumination. Error: %s", x, y, err.Error()))
				}
				colors = append(colors, resultColor)
			}
		}
	}

	return colors, nil
}

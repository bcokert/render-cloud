package raytracer

import (
	"github.com/bcokert/render-cloud/model"
	"github.com/bcokert/render-cloud/model/primitives"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"github.com/bcokert/render-cloud/raytracer/illumination/phong"
)

func FindRayCollision(origin, screenPoint mgl64.Vec3, objects []primitives.Primitive) (distance float64, primitive *primitives.Primitive) {
	ray := screenPoint.Sub(origin)

	for _, prim := range objects {
		collision := prim.FindClosestRayCollision(origin, ray)
		if collision != nil {
			if primitive == nil || *collision < distance {
				distance = *collision
			}
			primitive = &prim
		}
	}
	return distance, primitive
}

func FindRayColor(origin, screenPoint mgl64.Vec3, distance float64, primitive primitives.Primitive, ambientColor colorful.Color, lights []model.Light) colorful.Color {
	ray := screenPoint.Sub(origin).Normalize()
	viewerVector := ray.Mul(-1)
	normalVector, err := primitive.GetNormalAtPoint(ray.Mul(distance))
	if err != nil {
		log.Panicf("Failed to find the normal of object %v for origin %v, screenPoint %v, distance %#v: %s\n", primitive, origin, screenPoint, distance, err.Error())
	}

	resultColor := colorful.Color{0,0,0}
	for _, light := range lights {
		lightVector := light.Direction.Mul(-1).Normalize()
		reflectionVector := normalVector.Mul((2*lightVector.Dot(normalVector))).Normalize()

		phongColor, err := phong.IlluminateLocal(phong.MultiplyColors(ambientColor, primitive.GetMaterial().Color), *light.Color, primitive.GetMaterial().Color, lightVector, normalVector, viewerVector, reflectionVector, primitive.GetMaterial().Shininess)
		if err != nil {
			log.Panicf("Failed to illuminate with viewVector %v, normalVector %v, lightVector %v, reflectionVector %v. Error: %s", viewerVector, normalVector, lightVector, reflectionVector, err.Error())
		}

		resultColor = phong.CombineColors(resultColor, phongColor)
	}

	return resultColor
}

func TraceScene(scene model.Scene, width, height uint) (colors []colorful.Color) {
	ww, hh := int(width), int(height)
	screenWidth := 4.0
	screenHeight := 3.0
	screenDistance := 1.0

	screenForward := scene.World.Camera.Direction.Normalize()
	screenUp := scene.World.Camera.Up.Normalize()
	screenRight := screenUp.Cross(screenForward).Mul(-1).Normalize()

	screenCenter := scene.World.Camera.Origin.Add(screenForward.Mul(screenDistance))
	screenTopLeft := screenCenter.Add(screenUp.Mul(screenHeight / 2.0)).Sub(screenRight.Mul(screenWidth / 2.0))

	if len(scene.World.Lights) == 0 {
		log.Printf("There are no lights in the scene. This will make the whole scene the background color.")
	}

	for y := 0; y < hh; y++ {
		for x := 0; x < ww; x++ {

			// Optimization: no lights means collisions don't matter
			if len(scene.World.Lights) == 0 {
				append(colors, scene.World.Background)
				continue
			}

			// Create a ray through the screen at x,y and check for collisions
			dx := (float64(x) * screenWidth) / (float64(width))
			dy := (float64(y) * screenHeight) / (float64(height))
			screenPoint := screenTopLeft.Add(screenRight.Mul(dx)).Sub(screenUp.Mul(dy))
			distance, primitive := FindRayCollision(scene.World.Camera.Origin, screenPoint, *scene.Spheres)

			// If collisions, do local + global illumination. Else, return background color
			if primitive == nil {
				append(colors, scene.World.Background)
			} else {
				resultColor := FindRayColor(scene.World.Camera.Origin, screenPoint, distance, *primitive, scene.World.Ambient, *scene.World.Lights)
				append(colors, resultColor)
			}
		}
	}

	return colors
}

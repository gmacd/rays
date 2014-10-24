package testscenes

import (
	"github.com/gmacd/rays/core"
	"github.com/gmacd/rays/geom"
)

func CreateFlipcode3Scene() *geom.Scene {
	scene := geom.NewScene()

	// Ground plane
	plane := geom.NewPlane(core.NewVec3(0.0, 1.0, 0.0), 4.4)
	plane.SetName("Plane")
	plane.Material().Reflection = 0.0
	plane.Material().Refraction = 0.0
	plane.Material().Diffuse = 1.0
	plane.Material().Colour.Set(0.4, 0.3, 0.3)
	scene.AddPrimitive(plane)

	// Big sphere
	sphere := geom.NewSphere(core.NewVec3(2.0, 0.8, 3.0), 2.5)
	sphere.SetName("Big Sphere")
	sphere.Material().Reflection = 0.2
	sphere.Material().Refraction = 0.8
	sphere.Material().RefractiveIndex = 1.3
	sphere.Material().Colour.Set(0.7, 0.7, 1.0)
	scene.AddPrimitive(sphere)

	// Small sphere
	sphere = geom.NewSphere(core.NewVec3(-5.5, -0.5, 7.0), 2)
	sphere.SetName("Middling Sphere")
	sphere.Material().Reflection = 0.5
	sphere.Material().Refraction = 0.0
	sphere.Material().RefractiveIndex = 1.3
	sphere.Material().Diffuse = 0.1
	sphere.Material().Colour.Set(0.7, 0.7, 1.0)
	scene.AddPrimitive(sphere)

	// Light 1
	sphere = geom.NewSphere(core.NewVec3(0.0, 5.0, 5.0), 0.1)
	sphere.SetName("Wee Sphere (light)")
	sphere.SetIsLight(true)
	sphere.Material().Colour.Set(0.4, 0.4, 0.4)
	scene.AddPrimitive(sphere)

	// Light 2
	sphere = geom.NewSphere(core.NewVec3(-3.0, 5.0, 1.0), 0.1)
	sphere.SetName("Other light")
	sphere.SetIsLight(true)
	sphere.Material().Colour.Set(0.6, 0.6, 0.8)
	scene.AddPrimitive(sphere)

	// Extra sphere
	sphere = geom.NewSphere(core.NewVec3(-1.5, -3.8, 1.0), 1.5)
	sphere.SetName("Extra Sphere")
	sphere.Material().Reflection = 0.0
	sphere.Material().Refraction = 0.8
	sphere.Material().Colour.Set(1.0, 0.4, 0.4)
	scene.AddPrimitive(sphere)

	// Back plane
	plane = geom.NewPlane(core.NewVec3(0.4, 0.0, -1.0), 12)
	plane.SetName("Back Plane")
	plane.Material().Reflection = 0.0
	plane.Material().Refraction = 0.0
	plane.Material().Specular = 0.0
	plane.Material().Diffuse = 0.6
	plane.Material().Colour.Set(0.5, 0.3, 0.5)
	scene.AddPrimitive(plane)

	// Ceiling plane
	plane = geom.NewPlane(core.NewVec3(0, -1, 0), 7.4)
	plane.SetName("Ceiling Plane")
	plane.Material().Reflection = 0.0
	plane.Material().Refraction = 0.0
	plane.Material().Specular = 0.0
	plane.Material().Diffuse = 0.5
	plane.Material().Colour.Set(0.4, 0.7, 0.7)
	scene.AddPrimitive(plane)

	// Grid
	for x := 0.0; x < 8.0; x++ {
		for y := 0.0; y < 8.0; y++ {
			sphere = geom.NewSphere(core.NewVec3(-4.5+x*1.5, -4.3+y*1.5, 10), 0.3)
			sphere.SetName("Grid Sphere")
			sphere.Material().Reflection = 0.0
			sphere.Material().Refraction = 0.0
			sphere.Material().Specular = 0.6
			sphere.Material().Diffuse = 0.6
			sphere.Material().Colour.Set(0.3, 1.0, 0.4)
			scene.AddPrimitive(sphere)
		}
	}

	return scene
}

package testscenes

import (
	"fmt"
	"gmacd/core"
	"gmacd/geom"
	objParser "gmacd/modelparsers/obj"
	"gmacd/modelreaders/obj"
)

func CreateCubeScene() *geom.Scene {
	scene := geom.NewScene()

	objReader := objParser.NewObjReader(scene.Textures, "data/cube/cube.obj")
	if err := obj.Read(objReader); err != nil {
		fmt.Printf("** Error loading %v: %v\n", objReader.Filename(), err.Error())
		return scene
	}

	plane := geom.NewPlane(core.NewVec3(0.0, 1.0, 0.0), 4.4)
	plane.SetName("Plane")
	plane.Material().Reflection = 0.0
	plane.Material().Diffuse = 1.0
	plane.Material().Colour.Set(0.4, 0.3, 0.3)
	scene.AddPrimitive(plane)

	sphere := geom.NewSphere(core.NewVec3(1.0, -0.8, 3.0), 2.5)
	sphere.SetName("Big Sphere")
	sphere.Material().Reflection = 0.6
	sphere.Material().Colour.Set(0.7, 0.7, 0.7)
	scene.AddPrimitive(sphere)

	sphere = geom.NewSphere(core.NewVec3(-5.5, -0.5, 7.0), 2)
	sphere.SetName("Middling Sphere")
	sphere.Material().Reflection = 1.0
	sphere.Material().Diffuse = 0.1
	sphere.Material().Colour.Set(0.7, 0.7, 1.0)
	scene.AddPrimitive(sphere)

	sphere = geom.NewSphere(core.NewVec3(0.0, 5.0, 5.0), 0.1)
	sphere.SetName("Wee Sphere (light)")
	sphere.SetIsLight(true)
	sphere.Material().Colour.Set(0.6, 0.6, 0.6)
	scene.AddPrimitive(sphere)

	sphere = geom.NewSphere(core.NewVec3(2.0, 5.0, 1.0), 0.1)
	sphere.SetName("Other light")
	sphere.SetIsLight(true)
	sphere.Material().Colour.Set(0.7, 0.7, 0.9)
	scene.AddPrimitive(sphere)

	return scene
}

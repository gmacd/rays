package main

import (
	"bufio"
	"fmt"
	"gmacd/core"
	"gmacd/geom"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	canvas := NewCanvas(640, 480)

	render(canvas)

	if err := canvas.Export(); err != nil {
		fmt.Println(err.Error())
	}
}

type Canvas struct {
	img    *image.NRGBA
	width  int
	height int
}

func NewCanvas(width, height int) *Canvas {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	canvas := Canvas{img, width, height}
	return &canvas
}

func (canvas *Canvas) SetPixelRGB(x, y int, r, g, b uint8) {
	canvas.img.Set(x, y, color.NRGBA{r, g, b, 255})
}

//func (canvas *Canvas) SetPixelRGB(x, y int, r, g, b float64) {
//	canvas.img.Set(x, y, color.NRGBA{uint8(r * 255), uint8(g * 255), uint8(b * 255), 255})
//}

func (canvas *Canvas) SetPixel(x, y int, c core.ColourRGB) {
	canvas.img.Set(x, y, color.NRGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255})
}

func (canvas *Canvas) Export() error {
	imgFile, err := os.Create("xxx.png")
	if err != nil {
		return err
	}
	defer imgFile.Close()

	imgWriter := bufio.NewWriter(imgFile)
	if err = png.Encode(imgWriter, canvas.img); err != nil {
		return err
	}
	imgWriter.Flush()

	return nil
}

const MAX_TRACE_DEPTH = 6

func FindNearestIntersection(scene *Scene, ray core.Ray) (prim geom.Primitive, result int, dist float64) {
	dist = 1000000.0

	// Find nearest intersection
	for _, p := range scene.primitives {
		if pResult, pDist := p.Intersects(ray, dist); pResult != core.MISS {
			prim = p
			result = pResult
			dist = pDist
		}
	}

	return prim, result, dist
}

func Raytrace(scene *Scene, ray core.Ray, acc *core.ColourRGB, depth int, rIndex float64) (nearestPrim *geom.Primitive, dist float64) {
	if depth > MAX_TRACE_DEPTH {
		return nil, 0.0
	}

	prim, _, dist := FindNearestIntersection(scene, ray)
	if nearestPrim == nil {
		return nil, 0
	}

	if prim.IsLight() {
		acc.Set(1.0, 1.0, 1.0)
		return &prim, dist
	}

	// Determine intersection point
	intersectionPoint := ray.Origin.Add(ray.Dir.MulScalar(dist))

	// Trace lights
	for _, light := range scene.primitives {
		if !light.IsLight() {
			continue
		}

		// Calculate diffuse shading
		l := light.LightCentre().Sub(intersectionPoint).Normal()
		n := light.Normal(intersectionPoint)
		if prim.Material().Diffuse > 0 {
			dot := n.DotProduct(l)
			if dot > 0 {
				diff := dot * prim.Material().Diffuse
				acc.AddTo(prim.Material().Colour.MulScalar(diff).Mul(light.Material().Colour))
			}
		}
	}

	return &prim, dist
}

func render(canvas *Canvas) {
	scene := CreateScene()

	// init render
	wx1, wx2, wy1, wy2 := -4.0, 4.0, 3.0, -3.0
	sx, sy := 0.0, 0.0
	dx := (wx2 - wx1) / float64(canvas.width)
	dy := (wy2 - wy1) / float64(canvas.height)

	o := core.NewVec3(0, 0, -5)

	for y := 0; y < canvas.height; y++ {
		sx = wx1

		for x := 0; x < canvas.width; x++ {
			acc := core.NewColourRGB(0, 0, 0)
			dir := core.NewVec3(sx, sy, 0).Sub(o).Normal()
			ray := core.NewRay(o, dir)

			Raytrace(scene, ray, &acc, 1, 1.0)
			acc.R = math.Min(acc.R, 1.0)
			acc.G = math.Min(acc.G, 1.0)
			acc.B = math.Min(acc.B, 1.0)

			canvas.SetPixel(x, y, acc)

			sx += dx
		}
		sy += dy
	}
}

type Scene struct {
	primitives []geom.Primitive
}

func NewScene() *Scene {
	primitives := make([]geom.Primitive, 0, 20)
	return &Scene{primitives}
}

func (s *Scene) AddPrimitive(p geom.Primitive) {
	s.primitives = append(s.primitives, p)
}

// TODO Name<->Primitive map
func CreateScene() *Scene {
	scene := NewScene()

	plane := geom.NewPlane(core.NewVec3(0.0, 1.0, 0.0), 4.4)
	plane.Material().Reflection = 0.0
	plane.Material().Diffuse = 1.0
	plane.Material().Colour.Set(0.4, 0.3, 0.3)
	scene.AddPrimitive(plane)

	sphere := geom.NewSphere(core.NewVec3(1.0, -0.8, 3.0), 2.5)
	sphere.Material().Reflection = 0.6
	sphere.Material().Colour.Set(0.7, 0.7, 0.7)
	scene.AddPrimitive(sphere)

	sphere = geom.NewSphere(core.NewVec3(-5.5, -0.5, 7.0), 2)
	sphere.Material().Reflection = 1.0
	sphere.Material().Diffuse = 0.1
	sphere.Material().Colour.Set(0.7, 0.7, 1.0)
	scene.AddPrimitive(sphere)

	sphere = geom.NewSphere(core.NewVec3(0.0, 5.0, 5.0), 0.1)
	sphere.SetIsLight(true)
	sphere.Material().Colour.Set(0.6, 0.6, 0.6)
	scene.AddPrimitive(sphere)

	sphere = geom.NewSphere(core.NewVec3(2.0, 5.0, 1.0), 0.1)
	sphere.SetIsLight(true)
	sphere.Material().Colour.Set(0.7, 0.7, 0.9)
	scene.AddPrimitive(sphere)

	return scene
}

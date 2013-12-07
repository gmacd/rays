package main

import (
	"bufio"
	"fmt"
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

func (canvas *Canvas) SetPixel(x, y int, c ColourRGB) {
	canvas.img.Set(x, y, color.NRGBA{uint8(c.r * 255), uint8(c.g * 255), uint8(c.b * 255), 255})
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

func Raytrace(scene *Scene, ray Ray, acc *ColourRGB, depth int, rIndex float64) (nearestPrim *Primitive, dist float64) {
	if depth > MAX_TRACE_DEPTH {
		return nil, 0.0
	}

	dist = 1000000.0
	nearestPrim = nil
	//nearestIntersectionResult := MISS

	// Find nearest intersection
	for _, p := range scene.primitives {
		var result int
		if result, dist = p.Intersects(ray); result != MISS {
			nearestPrim = &p
			//nearestIntersectionResult = result
		}
	}

	if nearestPrim == nil {
		return nil, 0
	}

	// Eugh
	prim := *nearestPrim

	if prim.IsLight() {
		acc.Set(1.0, 1.0, 1.0)
		return &prim, dist
	}

	// Determine intersection point
	intersectionPoint := ray.origin.Add(ray.dir.MulScalar(dist))

	// Trace lights
	for _, light := range scene.primitives {
		if !light.IsLight() {
			continue
		}

		// Calculate diffuse shading
		// TODO remove dodgy cast...
		l := light.(Sphere).centre.Sub(intersectionPoint).Normal()
		n := light.Normal(intersectionPoint)
		if prim.Material().diffuse > 0 {
			dot := n.DotProduct(l)
			if dot > 0 {
				diff := dot * prim.Material().diffuse
				acc.AddTo(prim.Material().colour.MulScalar(diff).Mul(light.Material().colour))
			}
		}
	}

	return &prim, dist
}

func render(canvas *Canvas) {
	scene := NewScene()

	// init render
	wx1, wx2, wy1, wy2 := -4.0, 4.0, 3.0, -3.0
	sx, sy := 0.0, 0.0
	dx := (wx2 - wx1) / float64(canvas.width)
	dy := (wy2 - wy1) / float64(canvas.height)

	o := NewVec3(0, 0, -5)

	for y := 0; y < canvas.height; y++ {
		sx = wx1

		for x := 0; x < canvas.width; x++ {
			acc := NewColourRGB(0, 0, 0)
			dir := NewVec3(sx, sy, 0).Sub(o).Normal()
			ray := NewRay(o, dir)

			Raytrace(scene, ray, &acc, 1, 1.0)
			acc.r = math.Max(acc.r, 1.0)
			acc.g = math.Max(acc.g, 1.0)
			acc.b = math.Max(acc.b, 1.0)

			canvas.SetPixel(x, y, acc)

			sx += dx
		}
		sy += dy
	}
}

type Vec3 struct {
	x, y, z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(v1.x+v2.x, v1.y+v2.y, v1.z+v2.z)
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return NewVec3(v1.x-v2.x, v1.y-v2.y, v1.z-v2.z)
}

func (v1 Vec3) MulScalar(s float64) Vec3 {
	return NewVec3(v1.x*s, v1.y*s, v1.z*s)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v Vec3) Normalize() {
	l := v.Length()
	v.x /= l
	v.y /= l
	v.z /= l
}

func (v Vec3) Normal() Vec3 {
	l := v.Length()
	return NewVec3(v.x/l, v.y/l, v.z/l)
}

func (v1 Vec3) DotProduct(v2 Vec3) float64 {
	return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
}

type Ray struct {
	origin Vec3
	dir    Vec3
}

func NewRay(origin, dir Vec3) Ray {
	return Ray{origin, dir}
}

type ColourRGB struct {
	r, g, b float64
}

func NewColourRGB(r, g, b float64) ColourRGB {
	return ColourRGB{r, g, b}
}

func (c ColourRGB) Set(r, g, b float64) {
	c.r, c.g, c.b = r, g, b
}

func (c1 ColourRGB) AddTo(c2 ColourRGB) {
	c1.r, c1.g, c1.b = c2.r, c2.g, c2.b
}

func (c1 ColourRGB) Mul(c2 ColourRGB) ColourRGB {
	return NewColourRGB(c1.r*c2.r, c1.g*c2.g, c1.b*c2.b)
}

func (c ColourRGB) MulScalar(s float64) ColourRGB {
	return NewColourRGB(c.r*s, c.g*s, c.b*s)
}

const (
	MISS = iota
	HIT
	HIT_FROM_INSIDE
)

type Primitive interface {
	Intersects(ray Ray) (result int, dist float64)
	IsLight() bool
	Normal(v Vec3) Vec3
	Material() *Material
}

type Sphere struct {
	centre Vec3
	radius float64
	// TODO remove?  Premature?  Simplify?
	radiusSq    float64
	radiusRecip float64
	// TODO this seems wrong...
	isLight  bool
	material *Material
}

func NewSphere(centre Vec3, radius float64) *Sphere {
	return &Sphere{centre, radius, radius * radius, 1.0 / radius, false, nil}
}

func (sphere Sphere) Intersects(ray Ray) (result int, dist float64) {
	v := ray.origin.Sub(sphere.centre)
	b := -v.DotProduct(ray.dir)
	det := b*b - v.Length() + sphere.radiusSq

	if det > 0 {
		det = math.Sqrt(det)
		i2 := b + det

		if i2 > 0 {
			i1 := b - det

			if i1 < 0 {
				if i2 < dist {
					return HIT_FROM_INSIDE, i2
				}
			} else {
				if i1 < dist {
					return HIT, i1
				}
			}
		}
	}
	return MISS, 0
}

func (sphere Sphere) IsLight() bool {
	return sphere.isLight
}

func (sphere Sphere) Normal(v Vec3) Vec3 {
	return v.Sub(sphere.centre).MulScalar(sphere.radiusRecip)
}

func (sphere Sphere) Material() *Material {
	return sphere.material
}

type Scene struct {
	primitives []Primitive
}

func NewScene() *Scene {
	return &Scene{}
}

type Material struct {
	colour     ColourRGB
	reflection float64
	diffuse    float64
}

func NewMaterial(colour ColourRGB, reflection, diffuse float64) *Material {
	return &Material{colour, reflection, diffuse}
}

func (m Material) Specular() float64 {
	return 1.0 - m.diffuse
}

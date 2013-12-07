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

func Raytrace(ray Ray, acc *ColourRGB, depth int, rIndex float64) (p *Primitive, dist float64) {

}

func render(canvas *Canvas) {
	// init render
	wx1, wx2, wy1, wy2 := -4.0, 4.0, 3.0, -3.0
	sx, sy := 0.0, 0.0
	dx := (wx2 - wx1) / float64(canvas.width)
	dy := (wy2 - wy1) / float64(canvas.height)

	o := NewVec3(0, 0, -5)

	for y := 0; y < canvas.height; y++ {
		sx := wx1

		for x := 0; x < canvas.width; x++ {
			acc := NewColourRGB(0, 0, 0)
			dir := NewVec3(sx, sy, 0).Sub(o).Normal()
			ray := NewRay(o, dir)

			p, dist := Raytrace(ray, &acc, 1, 1.0)
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

type Primitive interface{}

type Sphere struct {
	centre                        Vec3
	radius, radiusSq, radiusRecip float64
}

func NewSphere(centre Vec3, radius float64) *Sphere {
	return &Sphere{centre, radius, radius * radius, 1.0 / radius}
}

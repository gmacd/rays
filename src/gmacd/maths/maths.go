package maths

import (
	"math"
)

type Vec3 struct {
	X, Y, Z float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func NewVec3Zero() Vec3 {
	return Vec3{0, 0, 0}
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return NewVec3(v1.X+v2.X, v1.Y+v2.Y, v1.Z+v2.Z)
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return NewVec3(v1.X-v2.X, v1.Y-v2.Y, v1.Z-v2.Z)
}

func (v1 Vec3) MulScalar(s float64) Vec3 {
	return NewVec3(v1.X*s, v1.Y*s, v1.Z*s)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Normalize() {
	l := v.Length()
	v.X /= l
	v.Y /= l
	v.Z /= l
}

func (v Vec3) Normal() Vec3 {
	l := v.Length()
	return NewVec3(v.X/l, v.Y/l, v.Z/l)
}

func (v1 Vec3) DotProduct(v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

type Ray struct {
	Origin Vec3
	Dir    Vec3
}

func NewRay(origin, dir Vec3) Ray {
	return Ray{origin, dir}
}

type Plane struct {
	Normal Vec3
	D      float64
}

func NewPlane(normal Vec3, d float64) *Plane {
	return &Plane{normal, d}
}

type ColourRGB struct {
	R, G, B float64
}

func NewColourRGB(r, g, b float64) ColourRGB {
	return ColourRGB{r, g, b}
}

func (c ColourRGB) Set(r, g, b float64) {
	c.R, c.G, c.B = r, g, b
}

func (c1 ColourRGB) AddTo(c2 ColourRGB) {
	c1.R, c1.G, c1.B = c2.R, c2.G, c2.B
}

func (c1 ColourRGB) Mul(c2 ColourRGB) ColourRGB {
	return NewColourRGB(c1.R*c2.R, c1.G*c2.G, c1.B*c2.B)
}

func (c ColourRGB) MulScalar(s float64) ColourRGB {
	return NewColourRGB(c.R*s, c.G*s, c.B*s)
}

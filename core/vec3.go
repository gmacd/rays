package core

import (
	"math"
)

// TODO Point

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

func (v1 Vec3) Mul(v2 Vec3) Vec3 {
	return NewVec3(v1.X*v2.X, v1.Y*v2.Y, v1.Z*v2.Z)
}

func (v1 Vec3) Div(v2 Vec3) Vec3 {
	return NewVec3(v1.X/v2.X, v1.Y/v2.Y, v1.Z/v2.Z)
}

func (v1 Vec3) MulScalar(s float64) Vec3 {
	return NewVec3(v1.X*s, v1.Y*s, v1.Z*s)
}

func (v1 Vec3) DivScalar(s float64) Vec3 {
	return NewVec3(v1.X/s, v1.Y/s, v1.Z/s)
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Normal() Vec3 {
	lr := 1.0 / v.Length()
	return NewVec3(v.X*lr, v.Y*lr, v.Z*lr)
}

func (v Vec3) NormalWithLength() (normal Vec3, originalLength float64) {
	l := v.Length()
	lr := 1.0 / l
	return NewVec3(v.X*lr, v.Y*lr, v.Z*lr), l
}

func (v1 Vec3) Dot(v2 Vec3) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z
}

func (v1 Vec3) Cross(v2 Vec3) (v3 Vec3) {
	return NewVec3(
		(v1.Y*v2.Z)-(v1.Z*v2.Y),
		(v1.Z*v2.X)-(v1.X*v2.Z),
		(v1.X*v2.Y)-(v1.Y*v2.X))
}

package core

import (
	"math"
)

// TODO Point

type Vec2 struct {
	X, Y float64
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

func NewVec2Zero() Vec2 {
	return Vec2{0, 0}
}

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return NewVec2(v1.X+v2.X, v1.Y+v2.Y)
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return NewVec2(v1.X-v2.X, v1.Y-v2.Y)
}

func (v1 Vec2) Mul(v2 Vec2) Vec2 {
	return NewVec2(v1.X*v2.X, v1.Y*v2.Y)
}

func (v1 Vec2) Div(v2 Vec2) Vec2 {
	return NewVec2(v1.X/v2.X, v1.Y/v2.Y)
}

func (v1 Vec2) MulScalar(s float64) Vec2 {
	return NewVec2(v1.X*s, v1.Y*s)
}

func (v1 Vec2) DivScalar(s float64) Vec2 {
	return NewVec2(v1.X/s, v1.Y/s)
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

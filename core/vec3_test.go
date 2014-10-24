package core

import (
	"math"
	"testing"
)

func TestNewVec3(t *testing.T) {
	v := NewVec3(1.0, 2.0, 3.0)
	if v.X != 1.0 || v.Y != 2.0 || v.Z != 3.0 {
		t.Error()
	}
}

func TestNewVec3Zero(t *testing.T) {
	v := NewVec3Zero()
	if v.X != 0.0 || v.Y != 0.0 || v.Z != 0.0 {
		t.Error()
	}
}

func TestVec3Add(t *testing.T) {
	v1 := NewVec3(1.0, 2.0, 3.0)
	v2 := NewVec3(10.0, 20.0, 30.0)
	v := v1.Add(v2)
	if v.X != 11.0 || v.Y != 22.0 || v.Z != 33.0 {
		t.Error()
	}
}

func TestVec3Sub(t *testing.T) {
	v1 := NewVec3(10.0, 20.0, 30.0)
	v2 := NewVec3(1.0, 2.0, 3.0)
	v := v1.Sub(v2)
	if v.X != 9.0 || v.Y != 18.0 || v.Z != 27.0 {
		t.Error()
	}
}

func TestVec3MulScalar(t *testing.T) {
	v1 := NewVec3(10.0, 20.0, 30.0)
	v := v1.MulScalar(10.0)
	if v.X != 100.0 || v.Y != 200.0 || v.Z != 300.0 {
		t.Error()
	}
}

func TestVec3Length(t *testing.T) {
	v := NewVec3(10.0, 10.0, 10.0)
	l := v.Length()
	if l != math.Sqrt(300.0) {
		t.Error()
	}
}

func TestVec3Normalize(t *testing.T) {
	v := NewVec3(1.0, 1.0, 1.0).Normal()
	a := 1.0 / math.Sqrt(3.0)
	if v.X != a || v.Y != a || v.Z != a {
		t.Error(v)
	}
}

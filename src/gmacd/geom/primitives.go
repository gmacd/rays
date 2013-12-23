package geom

import (
	"gmacd/core"
)

type Primitive interface {
	Name() string
	SetName(name string)

	Shape

	// TODO remove and become property of material?
	Light
}

type Shape interface {
	Intersects(ray core.Ray, maxDist float64) (result int, dist float64)
	Normal(v core.Vec3) core.Vec3
	Material() *core.Material
}

type PrimitiveData struct {
	name string
}

func NewPrimitiveData() *PrimitiveData {
	return &PrimitiveData{""}
}

func (primitive *PrimitiveData) Name() string {
	return primitive.name
}

func (primitive *PrimitiveData) SetName(name string) {
	primitive.name = name
}

package geom

import (
	"gmacd/core"
)

// TODO Lights as separate entities?
// TODO Light details as part of material?

type Primitive interface {
	Name() string
	SetName(name string)

	Shape
	Light
}

type Shape interface {
	Intersects(ray core.Ray, maxDist float64) (result int, dist float64)
	Normal(v core.Vec3) core.Vec3
	Material() *core.Material
}

type Light interface {
	IsLight() bool
	SetIsLight(isLight bool)
	LightCentre() core.Vec3
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

type LightData struct {
	isLight bool
	pos     core.Vec3
}

func NewLightData(pos core.Vec3) *LightData {
	return &LightData{false, pos}
}

func NewLightDataNone() *LightData {
	// TODO *************** Try nil
	// ****************************
	return &LightData{false, core.NewVec3Zero()}
}

func (light *LightData) IsLight() bool {
	return light.isLight
}

func (light *LightData) SetIsLight(isLight bool) {
	light.isLight = isLight
}

func (light *LightData) LightCentre() core.Vec3 {
	return light.pos
}

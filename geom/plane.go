package geom

import (
	"github.com/gmacd/rays/core"
	"github.com/gmacd/rays/intersections"
)

type Plane struct {
	plane *core.Plane

	material *core.Material

	*PrimitiveData
	*LightData
}

func NewPlane(normal core.Vec3, d float64) *Plane {
	p := core.NewPlane(normal, d)
	material := core.NewMaterialBlank()
	return &Plane{p, material, NewPrimitiveData(), NewLightDataNone()}
}

func (plane *Plane) Intersects(ray core.Ray, maxDist float64) (hit intersections.HitType, dist float64) {
	return intersections.IntersectRayPlane(ray, plane.plane, maxDist)
}

func (plane *Plane) Normal(v core.Vec3) core.Vec3 {
	return plane.plane.Normal
}

func (plane *Plane) Material() *core.Material {
	return plane.material
}

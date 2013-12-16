package geom

import (
	"gmacd/core"
)

type Plane struct {
	Plane core.Plane

	material *core.Material

	*PrimitiveData
	*LightData
}

func NewPlane(normal core.Vec3, d float64) *Plane {
	p := core.NewPlane(normal, d)
	material := core.NewMaterialBlank()
	return &Plane{*p, material, NewPrimitiveData(), NewLightDataNone()}
}

func (plane *Plane) Intersects(ray core.Ray, maxDist float64) (result int, dist float64) {
	d := plane.Plane.Normal.DotProduct(ray.Dir)
	if d == 0 {
		return core.MISS, 0.0
	}

	dist = -(plane.Plane.Normal.DotProduct(ray.Origin) + plane.Plane.D) / d
	if (dist > 0) && (dist < maxDist) {
		return core.HIT, dist
	}

	return core.MISS, 0.0
}

func (plane *Plane) Normal(v core.Vec3) core.Vec3 {
	return plane.Plane.Normal
}

func (plane *Plane) Material() *core.Material {
	return plane.material
}

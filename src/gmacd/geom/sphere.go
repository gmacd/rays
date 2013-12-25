package geom

import (
	"gmacd/core"
	"gmacd/intersections"
)

type Sphere struct {
	sphere   *core.Sphere
	material *core.Material

	*PrimitiveData
	*LightData
}

func NewSphere(centre core.Vec3, radius float64) *Sphere {
	material := core.NewMaterialBlank()
	return &Sphere{
		core.NewSphere(centre, radius), material,
		NewPrimitiveData(), NewLightData(centre)}
}

func (sphere *Sphere) Intersects(ray core.Ray, maxDist float64) intersections.HitDetails {
	return intersections.IntersectRaySphere(ray, sphere.sphere, maxDist)
}

func (sphere *Sphere) Normal(v core.Vec3) core.Vec3 {
	return v.Sub(sphere.sphere.Centre).MulScalar(sphere.sphere.RadiusRecip)
}

func (sphere *Sphere) Material() *core.Material {
	return sphere.material
}

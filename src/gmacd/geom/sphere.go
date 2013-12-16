package geom

import (
	"gmacd/core"
	"math"
)

type Sphere struct {
	centre      core.Vec3
	radius      float64
	radiusSq    float64
	radiusRecip float64

	material *core.Material

	*PrimitiveData
	*LightData
}

func NewSphere(centre core.Vec3, radius float64) *Sphere {
	material := core.NewMaterialBlank()
	return &Sphere{
		centre, radius, radius * radius, 1.0 / radius, material,
		NewPrimitiveData(), NewLightData(centre)}
}

func (sphere *Sphere) Intersects(ray core.Ray, maxDist float64) (result int, dist float64) {
	v := ray.Origin.Sub(sphere.centre)
	b := -v.DotProduct(ray.Dir)
	det := b*b - v.DotProduct(v) + sphere.radiusSq

	if det > 0 {
		det = math.Sqrt(det)
		i2 := b + det

		if i2 > 0 {
			i1 := b - det

			if i1 < 0 {
				if i2 < maxDist {
					return core.HIT_FROM_INSIDE, i2
				}
			} else {
				if i1 < maxDist {
					return core.HIT, i1
				}
			}
		}
	}
	return core.MISS, 0
}

func (sphere *Sphere) Normal(v core.Vec3) core.Vec3 {
	return v.Sub(sphere.centre).MulScalar(sphere.radiusRecip)
}

func (sphere *Sphere) Material() *core.Material {
	return sphere.material
}

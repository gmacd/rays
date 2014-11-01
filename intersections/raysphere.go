package intersections

import (
	"math"

	"github.com/gmacd/rays/core"
)

func IntersectRaySphere(r core.Ray, sphere *core.Sphere, maxDist float64) (hit HitType, dist float64) {
	v := r.Origin.Sub(sphere.Centre)
	b := -v.Dot(r.Dir)
	det := b*b - v.Dot(v) + sphere.RadiusSq

	if det > 0 {
		det = math.Sqrt(det)
		i2 := b + det

		if i2 > 0 {
			i1 := b - det

			if i1 < 0 {
				if i2 < maxDist {
					return HIT_FROM_INSIDE, i2
				}
			} else {
				if i1 < maxDist {
					return HIT, i1
				}
			}
		}
	}
	return MISS, 0
}

package intersections

import (
	"github.com/gmacd/rays/core"
)

func IntersectRayPlane(r core.Ray, plane *core.Plane, maxDist float64) (hit HitType, dist float64) {
	d := plane.Normal.Dot(r.Dir)
	if d == 0 {
		return MISS, 0
	}

	dist = -(plane.Normal.Dot(r.Origin) + plane.D) / d
	if (dist > 0) && (dist < maxDist) {
		return HIT, dist
	}

	return MISS, 0
}

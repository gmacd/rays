package intersections

import (
	"gmacd/core"
)

func IntersectRayPlane(r core.Ray, plane *core.Plane, maxDist float64) (result int, dist float64) {
	d := plane.Normal.Dot(r.Dir)
	if d == 0 {
		return core.MISS, 0.0
	}

	dist = -(plane.Normal.Dot(r.Origin) + plane.D) / d
	if (dist > 0) && (dist < maxDist) {
		return core.HIT, dist
	}

	return core.MISS, 0.0
}

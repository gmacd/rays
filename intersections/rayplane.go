package intersections

import (
	"github.com/gmacd/rays/core"
)

func IntersectRayPlane(r core.Ray, plane *core.Plane, maxDist float64) HitDetails {
	d := plane.Normal.Dot(r.Dir)
	if d == 0 {
		return NewMiss()
	}

	dist := -(plane.Normal.Dot(r.Origin) + plane.D) / d
	if (dist > 0) && (dist < maxDist) {
		return NewHit(dist)
	}

	return NewMiss()
}

package intersections

import (
	"github.com/gmacd/rays/core"
)

// Assumes the ray and triangle points are in the same space
func IntersectRayTriangle(r core.Ray, p1, p2, p3 core.Vec3, maxDist float64) (hit HitType, dist float64) {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)
	s1 := r.Dir.Cross(e2)
	divisor := s1.Dot(e1)
	if divisor == 0 {
		return MISS, 0
	}
	invDivisor := 1.0 / divisor

	// First barycentric coord
	d := r.Origin.Sub(p1)
	b1 := d.Dot(s1) * invDivisor
	if b1 < 0 || b1 > 1 {
		return MISS, 0
	}

	// Second barycentric coord
	s2 := d.Cross(e1)
	b2 := r.Dir.Dot(s2) * invDivisor
	if b2 < 0 || b2 > 1 {
		return MISS, 0
	}

	// We have an intersection, but early out if it's past maxDist
	t := e2.Dot(s2) * invDivisor
	if t < 0 || t > maxDist {
		return MISS, 0
	}

	return HIT, t
}

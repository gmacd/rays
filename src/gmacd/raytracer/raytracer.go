package raytracer

import (
	"gmacd/core"
	"gmacd/geom"
)

const MAX_TRACE_DEPTH = 6

func FindNearestIntersection(scene *geom.Scene, ray core.Ray) (prim geom.Primitive, result int, dist float64) {
	dist = 1000000.0

	// Find nearest intersection
	result = core.MISS
	for _, p := range scene.Primitives {
		if pResult, pDist := p.Intersects(ray, dist); pResult != core.MISS {
			prim = p
			result = pResult
			dist = pDist
		}
	}

	return prim, result, dist
}

func Raytrace(scene *geom.Scene, ray core.Ray, acc *core.ColourRGB, depth int, rIndex float64) (prim geom.Primitive, dist float64) {
	if depth > MAX_TRACE_DEPTH {
		return nil, 0.0
	}

	prim, _, dist = FindNearestIntersection(scene, ray)
	if prim == nil {
		return nil, 0
	}

	if prim.IsLight() {
		acc.Set(1.0, 1.0, 1.0)
		return prim, dist
	}

	// Determine intersection point
	intersectionPoint := ray.Origin.Add(ray.Dir.MulScalar(dist))

	// Trace lights
	for _, light := range scene.Primitives {
		if !light.IsLight() {
			continue
		}

		// Calculate diffuse shading
		l := light.LightCentre().Sub(intersectionPoint).Normal()
		n := prim.Normal(intersectionPoint)
		if prim.Material().Diffuse > 0 {
			dot := n.DotProduct(l)
			if dot > 0 {
				diff := dot * prim.Material().Diffuse
				acc.AddTo(prim.Material().Colour.MulScalar(diff).Mul(light.Material().Colour))
			}
		}
	}

	// Calculate reflection
	reflection := prim.Material().Reflection
	if reflection > 0 {
		n := prim.Normal(intersectionPoint)
		r := ray.Dir.Sub(n.MulScalar(2.0 * ray.Dir.DotProduct(n)))
		if depth < MAX_TRACE_DEPTH {
			reflectionRay := core.NewRay(
				intersectionPoint.Add(r.MulScalar(core.EPSILON)), r)
			reflectionColour := core.NewColourRGB(0, 0, 0)

			Raytrace(scene, reflectionRay, &reflectionColour, depth+1, rIndex)
			acc.AddTo(prim.Material().Colour.Mul(reflectionColour).MulScalar(reflection))
		}
	}

	return prim, dist
}

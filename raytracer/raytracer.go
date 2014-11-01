package raytracer

import (
	"math"

	"github.com/gmacd/rays/core"
	"github.com/gmacd/rays/geom"
	"github.com/gmacd/rays/intersections"
)

const MAX_TRACE_DEPTH = 5

func FindNearestIntersection(scene *geom.Scene, ray core.Ray) (hitPrim geom.Primitive, hit intersections.HitType, dist float64) {
	maxDist := math.MaxFloat64

	// Find nearest intersection
	for _, p := range scene.AllPrimitives() {
		currHit, currHitDist := p.Intersects(ray, maxDist)
		if currHit != intersections.MISS {
			hitPrim = p
			hit = currHit
			dist = currHitDist
		}
	}

	return hitPrim, hit, dist
}

func Raytrace(scene *geom.Scene, ray core.Ray, rIndex float64) (hitPrim geom.Primitive, hit intersections.HitType, dist float64, colour core.ColourRGB) {
	colour = core.NewColourRGB(0, 0, 0)
	hitPrim, hit, dist = FindNearestIntersection(scene, ray)
	if hit == intersections.MISS {
		return hitPrim, hit, dist, colour
	}

	// This is a bit rubbish - always white for direct light hit?
	if hitPrim.IsLight() {
		return hitPrim, hit, dist, core.NewColourRGB(1, 1, 1)
	}

	material := hitPrim.Material()
	intersectionPoint := ray.Origin.Add(ray.Dir.MulScalar(dist))

	// Trace lights
	for _, light := range scene.Lights() {
		// Point light shadows
		shade := 1.0
		// TODO is point light?
		l, lightDist := light.LightCentre().Sub(intersectionPoint).NormalWithLength()
		{
			// If point light
			r := core.NewRayWithDepth(intersectionPoint.Add(l.MulScalar(core.EPSILON)), l, ray.Depth+1)
			for _, primForShadow := range scene.AllPrimitives() {
				if primForShadow != light {
					if lightHit, _ := primForShadow.Intersects(r, lightDist); lightHit != intersections.MISS {
						shade = 0
						break
					}
				}
			}
		}

		// Diffuse
		n := hitPrim.Normal(intersectionPoint)
		if material.Diffuse > 0 {
			dot := n.Dot(l)
			if dot > 0 {
				diffuse := dot * material.Diffuse * shade
				colour.AddTo(material.Colour.MulScalar(diffuse).Mul(light.Material().Colour))
			}
		}

		// Specular
		specular := material.Specular
		v := ray.Dir
		r := l.Sub(n.MulScalar(2.0 * l.Dot(n)))
		dot := v.Dot(r)
		if dot > 0 {
			specular := math.Pow(dot, 20) * specular * shade
			colour.AddTo(light.Material().Colour.MulScalar(specular))
		}
	}

	// Reflection
	reflection := material.Reflection
	if (reflection > 0) && (ray.Depth < MAX_TRACE_DEPTH) {
		n := hitPrim.Normal(intersectionPoint)
		r := ray.Dir.Sub(n.MulScalar(2.0 * ray.Dir.Dot(n)))
		reflectionRay := core.NewRayWithDepth(
			intersectionPoint.Add(r.MulScalar(core.EPSILON)), r, ray.Depth+1)

		_, _, _, reflectionColour := Raytrace(scene, reflectionRay, rIndex)
		colour.AddTo(material.Colour.Mul(reflectionColour).MulScalar(reflection))
	}

	// Refraction
	refraction := material.Refraction
	if (refraction > 0) && (ray.Depth < MAX_TRACE_DEPTH) {
		primRIndex := material.RefractiveIndex
		n := rIndex / primRIndex
		N := hitPrim.Normal(intersectionPoint).MulScalar(float64(hit))
		cosI := -N.Dot(ray.Dir)
		cosT2 := 1.0 - n*n*(1.0-cosI*cosI)
		if cosT2 > 0 {
			T := ray.Dir.MulScalar(n).Add(N.MulScalar(n*cosI - math.Sqrt(cosT2)))
			refractiveRay := core.NewRayWithDepth(
				intersectionPoint.Add(T.MulScalar(core.EPSILON)),
				T, ray.Depth+1)
			_, _, refractiveHitDist, refractiveColour := Raytrace(scene, refractiveRay, primRIndex)

			absorbance := material.Colour.MulScalar(0.15 * -refractiveHitDist)
			transparency := core.NewColourRGB(
				math.Exp(absorbance.R),
				math.Exp(absorbance.G),
				math.Exp(absorbance.B))

			colour.AddTo(refractiveColour.Mul(transparency))
		}
	}

	return hitPrim, hit, dist, colour
}

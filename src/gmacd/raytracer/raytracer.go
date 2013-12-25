package raytracer

import (
	"gmacd/core"
	"gmacd/geom"
	"gmacd/intersections"
	"math"
)

const MAX_TRACE_DEPTH = 5

func FindNearestIntersection(scene *geom.Scene, ray core.Ray) (hitPrim geom.Primitive, hitDetails intersections.HitDetails) {
	maxDist := math.MaxFloat64

	// Find nearest intersection
	hitDetails = intersections.NewMiss()
	for _, p := range scene.AllPrimitives() {
		if currHitDetails := p.Intersects(ray, maxDist); currHitDetails.IsAnyHit() {
			hitPrim = p
			hitDetails = currHitDetails
			maxDist = currHitDetails.Dist
		}
	}

	return hitPrim, hitDetails
}

func Raytrace(scene *geom.Scene, ray core.Ray, rIndex float64) (hitPrim geom.Primitive, hitDetails intersections.HitDetails, colour core.ColourRGB) {
	colour = core.NewColourRGB(0, 0, 0)
	hitPrim, hitDetails = FindNearestIntersection(scene, ray)
	if hitDetails.IsMiss() {
		return hitPrim, hitDetails, colour
	}

	// This is a bit rubbish - always white for direct light hit?
	if hitPrim.IsLight() {
		return hitPrim, hitDetails, core.NewColourRGB(1, 1, 1)
	}

	material := hitPrim.Material()
	intersectionPoint := ray.Origin.Add(ray.Dir.MulScalar(hitDetails.Dist))

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
					if primForShadow.Intersects(r, lightDist).IsAnyHit() {
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

		_, _, reflectionColour := Raytrace(scene, reflectionRay, rIndex)
		colour.AddTo(material.Colour.Mul(reflectionColour).MulScalar(reflection))
	}

	// Refraction
	refraction := material.Refraction
	if (refraction > 0) && (ray.Depth < MAX_TRACE_DEPTH) {
		primRIndex := material.RefractiveIndex
		n := rIndex / primRIndex
		N := hitPrim.Normal(intersectionPoint).MulScalar(float64(hitDetails.Result))
		cosI := -N.Dot(ray.Dir)
		cosT2 := 1.0 - n*n*(1.0-cosI*cosI)
		if cosT2 > 0 {
			T := ray.Dir.MulScalar(n).Add(N.MulScalar(n*cosI - math.Sqrt(cosT2)))
			refractiveRay := core.NewRayWithDepth(
				intersectionPoint.Add(T.MulScalar(core.EPSILON)),
				T, ray.Depth+1)
			_, refractiveHitDetails, refractiveColour := Raytrace(scene, refractiveRay, primRIndex)

			absorbance := material.Colour.MulScalar(0.15 * -refractiveHitDetails.Dist)
			transparency := core.NewColourRGB(
				math.Exp(absorbance.R),
				math.Exp(absorbance.G),
				math.Exp(absorbance.B))

			colour.AddTo(refractiveColour.Mul(transparency))
		}
	}

	return hitPrim, hitDetails, colour
}

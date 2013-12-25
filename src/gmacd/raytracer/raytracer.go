package raytracer

import (
	"gmacd/core"
	"gmacd/geom"
	"gmacd/intersections"
	"math"
)

const MAX_TRACE_DEPTH = 5

func FindNearestIntersection(scene *geom.Scene, ray core.Ray) (prim geom.Primitive, hitDetails intersections.HitDetails) {
	maxDist := 1000000.0

	// Find nearest intersection
	hitDetails = intersections.NewMiss()
	for _, p := range scene.Primitives {
		if currHitDetails := p.Intersects(ray, maxDist); currHitDetails.IsAnyHit() {
			prim = p
			hitDetails = currHitDetails
			maxDist = currHitDetails.Dist
		}
	}

	return prim, hitDetails
}

func Raytrace(scene *geom.Scene, ray core.Ray, acc *core.ColourRGB, rIndex float64) (prim geom.Primitive, dist float64) {
	prim, hitDetails := FindNearestIntersection(scene, ray)
	if prim == nil {
		return nil, 0
	}
	dist = hitDetails.Dist

	// This is a bit rubbish - always white for direct light hit?
	if prim.IsLight() {
		acc.Set(1.0, 1.0, 1.0)
		return prim, dist
	}

	intersectionPoint := ray.Origin.Add(ray.Dir.MulScalar(dist))

	// Trace lights
	for _, light := range scene.Primitives {
		if !light.IsLight() {
			continue
		}

		// Point light shadows
		shade := 1.0
		// TODO is point light?
		l, lightDist := light.LightCentre().Sub(intersectionPoint).NormalWithLength()
		{
			// If point light
			r := core.NewRayWithDepth(intersectionPoint.Add(l.MulScalar(core.EPSILON)), l, ray.Depth+1)
			for _, primForShadow := range scene.Primitives {
				if primForShadow != light {
					if result := primForShadow.Intersects(r, lightDist); result.IsAnyHit() {
						shade = 0
						break
					}
				}
			}
		}

		// Diffuse
		n := prim.Normal(intersectionPoint)
		if prim.Material().Diffuse > 0 {
			dot := n.Dot(l)
			if dot > 0 {
				diffuse := dot * prim.Material().Diffuse * shade
				acc.AddTo(prim.Material().Colour.MulScalar(diffuse).Mul(light.Material().Colour))
			}
		}

		// Specular
		specular := prim.Material().Specular
		v := ray.Dir
		r := l.Sub(n.MulScalar(2.0 * l.Dot(n)))
		dot := v.Dot(r)
		if dot > 0 {
			specular := math.Pow(dot, 20) * specular * shade
			acc.AddTo(light.Material().Colour.MulScalar(specular))
		}
	}

	// Reflection
	reflection := prim.Material().Reflection
	if (reflection > 0) && (ray.Depth < MAX_TRACE_DEPTH) {
		n := prim.Normal(intersectionPoint)
		r := ray.Dir.Sub(n.MulScalar(2.0 * ray.Dir.Dot(n)))
		reflectionRay := core.NewRayWithDepth(
			intersectionPoint.Add(r.MulScalar(core.EPSILON)), r, ray.Depth+1)
		reflectionColour := core.NewColourRGB(0, 0, 0)

		Raytrace(scene, reflectionRay, &reflectionColour, rIndex)
		acc.AddTo(prim.Material().Colour.Mul(reflectionColour).MulScalar(reflection))
	}

	// Refraction
	refraction := prim.Material().Refraction
	if (refraction > 0) && (ray.Depth < MAX_TRACE_DEPTH) {
		primRIndex := prim.Material().RefractiveIndex
		n := rIndex / primRIndex
		N := prim.Normal(intersectionPoint).MulScalar(float64(hitDetails.Result))
		cosI := -N.Dot(ray.Dir)
		cosT2 := 1.0 - n*n*(1.0-cosI*cosI)
		if cosT2 > 0 {
			T := ray.Dir.MulScalar(n).Add(N.MulScalar(n*cosI - math.Sqrt(cosT2)))
			refractiveColour := core.NewColourRGB(0, 0, 0)
			refractiveRay := core.NewRayWithDepth(
				intersectionPoint.Add(T.MulScalar(core.EPSILON)),
				T, ray.Depth+1)
			_, refractiveDist := Raytrace(scene, refractiveRay, &refractiveColour, primRIndex)

			absorbance := prim.Material().Colour.MulScalar(0.15 * -refractiveDist)
			transparency := core.NewColourRGB(
				math.Exp(absorbance.R),
				math.Exp(absorbance.G),
				math.Exp(absorbance.B))

			acc.AddTo(refractiveColour.Mul(transparency))
		}
	}

	return prim, dist
}

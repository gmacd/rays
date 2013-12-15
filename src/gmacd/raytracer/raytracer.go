package raytracer

import (
	"gmacd/core"
	"gmacd/geom"
	"math"
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

	intersectionResult := core.MISS
	prim, intersectionResult, dist = FindNearestIntersection(scene, ray)
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

		// Point light shadows
		shade := 1.0
		// TODO is point light?
		l, lightDist := light.LightCentre().Sub(intersectionPoint).NormalWithLength()
		{
			// If point light
			r := core.NewRay(intersectionPoint.Add(l.MulScalar(core.EPSILON)), l)
			for _, primForShadow := range scene.Primitives {
				if primForShadow != light {
					if result, _ := primForShadow.Intersects(r, lightDist); result != core.MISS {
						shade = 0
						break
					}
				}
			}
		}

		// Calculate diffuse shading
		n := prim.Normal(intersectionPoint)
		if prim.Material().Diffuse > 0 {
			dot := n.DotProduct(l)
			if dot > 0 {
				diffuse := dot * prim.Material().Diffuse * shade
				acc.AddTo(prim.Material().Colour.MulScalar(diffuse).Mul(light.Material().Colour))
			}
		}

		// Specular
		specular := prim.Material().Specular
		v := ray.Dir
		r := l.Sub(n.MulScalar(2.0 * l.DotProduct(n)))
		dot := v.DotProduct(r)
		if dot > 0 {
			specular := math.Pow(dot, 20) * specular * shade
			acc.AddTo(light.Material().Colour.MulScalar(specular))
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

	// Refraction
	refraction := prim.Material().Refraction
	if refraction > 0 && depth < MAX_TRACE_DEPTH {
		primRIndex := prim.Material().RefractiveIndex
		n := rIndex / primRIndex
		N := prim.Normal(intersectionPoint).MulScalar(float64(intersectionResult))
		cosI := -N.DotProduct(ray.Dir)
		cosT2 := 1.0 - n*n*(1.0-cosI*cosI)
		if cosT2 > 0 {
			T := ray.Dir.MulScalar(n).Add(N.MulScalar(n*cosI - math.Sqrt(cosT2)))
			refractiveColour := core.NewColourRGB(0, 0, 0)
			refractiveRay := core.NewRay(
				intersectionPoint.Add(T.MulScalar(core.EPSILON)),
				T)
			_, refractiveDist := Raytrace(scene, refractiveRay, &refractiveColour, depth+1, primRIndex)

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

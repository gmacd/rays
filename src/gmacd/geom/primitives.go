package geom

import (
	"gmacd/core"
	"gmacd/maths"
	"math"
)

type Primitive interface {
	Intersects(ray maths.Ray, maxDist float64) (result int, dist float64)
	Normal(v maths.Vec3) maths.Vec3
	Material() *Material
	IsLight() bool
	SetIsLight(isLight bool)
	LightCentre() maths.Vec3
}

type Sphere struct {
	Centre maths.Vec3
	Radius float64
	// TODO remove?  Premature?  Simplify?
	RadiusSq    float64
	RadiusRecip float64
	// TODO this seems wrong...
	isLight  bool
	material *Material
}

func NewSphere(centre maths.Vec3, radius float64) *Sphere {
	material := NewMaterialBlank()
	return &Sphere{centre, radius, radius * radius, 1.0 / radius, false, material}
}

func (sphere *Sphere) Intersects(ray maths.Ray, maxDist float64) (result int, dist float64) {
	v := ray.Origin.Sub(sphere.Centre)
	b := -v.DotProduct(ray.Dir)
	det := b*b - v.Length() + sphere.RadiusSq

	if det > 0 {
		det = math.Sqrt(det)
		i2 := b + det

		if i2 > 0 {
			i1 := b - det

			if i1 < 0 {
				if i2 < maxDist {
					return core.HIT_FROM_INSIDE, i2
				}
			} else {
				if i1 < maxDist {
					return core.HIT, i1
				}
			}
		}
	}
	return core.MISS, 0
}

func (sphere *Sphere) IsLight() bool {
	return sphere.isLight
}

func (sphere *Sphere) SetIsLight(isLight bool) {
	sphere.isLight = isLight
}

func (sphere *Sphere) Normal(v maths.Vec3) maths.Vec3 {
	return v.Sub(sphere.Centre).MulScalar(sphere.RadiusRecip)
}

func (sphere *Sphere) Material() *Material {
	return sphere.material
}

func (sphere *Sphere) LightCentre() maths.Vec3 {
	return sphere.Centre
}

type Plane struct {
	Plane    maths.Plane
	material *Material
}

func NewPlane(normal maths.Vec3, d float64) *Plane {
	p := maths.NewPlane(normal, d)
	material := NewMaterialBlank()
	return &Plane{*p, material}
}

func (plane *Plane) Intersects(ray maths.Ray, maxDist float64) (result int, dist float64) {
	d := plane.Plane.Normal.DotProduct(ray.Dir)
	if d == 0 {
		return core.MISS, 0.0
	}

	dist = -(plane.Plane.Normal.DotProduct(ray.Origin) + plane.Plane.D) / d
	if (dist > 0) && (dist < maxDist) {
		return core.HIT, dist
	}

	return core.MISS, 0.0
}

func (plane *Plane) IsLight() bool {
	return false
}

func (plane *Plane) SetIsLight(isLight bool) {}

func (plane *Plane) LightCentre() maths.Vec3 {
	return maths.NewVec3Zero()
}

func (plane *Plane) Normal(v maths.Vec3) maths.Vec3 {
	return plane.Plane.Normal
}

func (plane *Plane) Material() *Material {
	return plane.material
}

type Material struct {
	Colour     maths.ColourRGB
	Reflection float64
	Diffuse    float64
}

func NewMaterialBlank() *Material {
	return &Material{maths.NewColourRGB(0.0, 0.0, 0.0), 0.0, 0.0}
}

func NewMaterial(colour maths.ColourRGB, reflection, diffuse float64) *Material {
	return &Material{colour, reflection, diffuse}
}

func (m *Material) Specular() float64 {
	return 1.0 - m.Diffuse
}

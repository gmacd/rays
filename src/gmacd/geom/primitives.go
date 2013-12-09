package geom

import (
	"gmacd/core"
	"math"
)

type Primitive interface {
	Intersects(ray core.Ray, maxDist float64) (result int, dist float64)
	Normal(v core.Vec3) core.Vec3
	Material() *Material
	IsLight() bool
	SetIsLight(isLight bool)
	LightCentre() core.Vec3
}

type Sphere struct {
	Centre core.Vec3
	Radius float64
	// TODO remove?  Premature?  Simplify?
	RadiusSq    float64
	RadiusRecip float64
	// TODO this seems wrong...
	isLight  bool
	material *Material
}

func NewSphere(centre core.Vec3, radius float64) *Sphere {
	material := NewMaterialBlank()
	return &Sphere{centre, radius, radius * radius, 1.0 / radius, false, material}
}

func (sphere *Sphere) Intersects(ray core.Ray, maxDist float64) (result int, dist float64) {
	v := ray.Origin.Sub(sphere.Centre)
	b := -v.DotProduct(ray.Dir)
	det := b*b - v.Length() + sphere.RadiusSq

	if det > 0 {
		det = math.Sqrt(det)
		i2 := b + det

		if i2 > 0 {
			i1 := b - det

			//fmt.Printf("i1=%v, i2=%v\n", i1, i2)

			if i1 < 0 {
				if i2 < maxDist {
					//fmt.Println("HIT_FROM_INSIDE")
					return core.HIT_FROM_INSIDE, i2
				}
			} else {
				if i1 < maxDist {
					//fmt.Println("HIT")
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

func (sphere *Sphere) Normal(v core.Vec3) core.Vec3 {
	return v.Sub(sphere.Centre).MulScalar(sphere.RadiusRecip)
}

func (sphere *Sphere) Material() *Material {
	return sphere.material
}

func (sphere *Sphere) LightCentre() core.Vec3 {
	return sphere.Centre
}

type Plane struct {
	Plane    core.Plane
	material *Material
}

func NewPlane(normal core.Vec3, d float64) *Plane {
	p := core.NewPlane(normal, d)
	material := NewMaterialBlank()
	return &Plane{*p, material}
}

func (plane *Plane) Intersects(ray core.Ray, maxDist float64) (result int, dist float64) {
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

func (plane *Plane) LightCentre() core.Vec3 {
	return core.NewVec3Zero()
}

func (plane *Plane) Normal(v core.Vec3) core.Vec3 {
	return plane.Plane.Normal
}

func (plane *Plane) Material() *Material {
	return plane.material
}

type Material struct {
	Colour     core.ColourRGB
	Reflection float64
	Diffuse    float64
}

func NewMaterialBlank() *Material {
	return &Material{core.NewColourRGB(0.0, 0.0, 0.0), 0.0, 0.0}
}

func NewMaterial(colour core.ColourRGB, reflection, diffuse float64) *Material {
	return &Material{colour, reflection, diffuse}
}

func (m Material) Specular() float64 {
	return 1.0 - m.Diffuse
}

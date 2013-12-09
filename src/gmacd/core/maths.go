package core

type Ray struct {
	Origin Vec3
	Dir    Vec3
}

func NewRay(origin, dir Vec3) Ray {
	return Ray{origin, dir}
}

type Plane struct {
	Normal Vec3
	D      float64
}

func NewPlane(normal Vec3, d float64) *Plane {
	return &Plane{normal, d}
}

type ColourRGB struct {
	R, G, B float64
}

func NewColourRGB(r, g, b float64) ColourRGB {
	return ColourRGB{r, g, b}
}

func (c *ColourRGB) Set(r, g, b float64) {
	c.R, c.G, c.B = r, g, b
}

func (c1 *ColourRGB) AddTo(c2 ColourRGB) {
	c1.R, c1.G, c1.B = c2.R, c2.G, c2.B
}

func (c1 ColourRGB) Mul(c2 ColourRGB) ColourRGB {
	return NewColourRGB(c1.R*c2.R, c1.G*c2.G, c1.B*c2.B)
}

func (c ColourRGB) MulScalar(s float64) ColourRGB {
	return NewColourRGB(c.R*s, c.G*s, c.B*s)
}

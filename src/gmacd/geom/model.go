package geom

import (
	"gmacd/core"
	"gmacd/intersections"
)

type Tri struct {
	Vertices        []*core.Vec3
	TextureVertices []*core.Vec2
	Normals         []*core.Vec3
}

func NewTri() *Tri {
	return &Tri{
		make([]*core.Vec3, 0, 3),
		make([]*core.Vec2, 0, 3),
		make([]*core.Vec3, 0, 3)}
}

type Model struct {
	Vertices        []core.Vec3
	TextureVertices []core.Vec2
	Normals         []core.Vec3
	Triangles       []Tri

	Materials map[string]*core.Material

	*PrimitiveData
	*LightData
}

func NewModel() *Model {
	return &Model{
		make([]core.Vec3, 0, 100),
		make([]core.Vec2, 0, 100),
		make([]core.Vec3, 0, 100),
		make([]Tri, 0, 100),
		make(map[string]*core.Material),

		NewPrimitiveData(),
		NewLightDataNone(),
	}
}

func (model *Model) Intersects(ray core.Ray, maxDist float64) intersections.HitDetails {
	/*for _, tri := range model.Triangles {
		hit, dist := intersections.IntersectRayTriangle(ray, tri.Vertices[0], tri.Vertices[1], tri.Vertices[2], maxDist)
		if hit {
			return core.HIT, dist
		}
	}*/
	return intersections.NewMiss()
}

/*func (model *Model) Normal(v core.Vec3) core.Vec3 {
	return plane.plane.Normal
}

func (model *Model) Material() *core.Material {
	return plane.material
}*/

func NewSingleTriangleModel(p1, p2, p3 core.Vec3) *Model {
	model := NewModel()
	model.Vertices = append(model.Vertices, p1, p2, p3)

	tri := NewTri()
	tri.Vertices = append(tri.Vertices, &model.Vertices[0], &model.Vertices[1], &model.Vertices[2])

	model.Triangles = append(model.Triangles)
	return model
}

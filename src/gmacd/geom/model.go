package geom

import (
	"gmacd/core"
)

type Model struct {
	Vertices        []core.Vec3
	TextureVertices []core.Vec2
	Normals         []core.Vec3
	Triangles       []Tri

	Materials map[string]*core.Material
}

func NewModel() *Model {
	return &Model{
		make([]core.Vec3, 0, 100),
		make([]core.Vec2, 0, 100),
		make([]core.Vec3, 0, 100),
		make([]Tri, 0, 100),
		make(map[string]*core.Material),
	}
}

type Tri struct {
	Vertex        *core.Vec3
	TextureVertex *core.Vec2
	Normal        *core.Vec3
}

func NewTri() *Tri {
	return &Tri{nil, nil, nil}
}

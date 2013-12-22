package geom

import (
	"gmacd/core"
)

type Scene struct {
	Primitives []Primitive
	Textures   core.Textures
}

func NewScene() *Scene {
	primitives := make([]Primitive, 0, 20)
	return &Scene{primitives, core.NewTextures()}
}

func (s *Scene) AddPrimitive(p Primitive) {
	s.Primitives = append(s.Primitives, p)
}

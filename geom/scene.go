package geom

import (
	"github.com/gmacd/rays/core"
)

type Scene struct {
	allPrimitives []Primitive
	lights        []Primitive
	textures      core.Textures
}

func NewScene() *Scene {
	return &Scene{
		make([]Primitive, 0, 100),
		make([]Primitive, 0, 100),
		core.NewTextures()}
}

func (s *Scene) AddPrimitive(p Primitive) {
	s.allPrimitives = append(s.allPrimitives, p)
	if p.IsLight() {
		s.lights = append(s.lights, p)
	}
}

func (s *Scene) AllPrimitives() []Primitive {
	return s.allPrimitives
}

func (s *Scene) Lights() []Primitive {
	return s.lights
}

func (s *Scene) Textures() core.Textures {
	return s.textures
}

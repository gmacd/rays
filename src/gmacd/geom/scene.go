package geom

type Scene struct {
	Primitives []Primitive
}

func NewScene() *Scene {
	primitives := make([]Primitive, 0, 20)
	return &Scene{primitives}
}

func (s *Scene) AddPrimitive(p Primitive) {
	s.Primitives = append(s.Primitives, p)
}

package core

type Ray struct {
	Origin Vec3
	Dir    Vec3
	Depth  int
}

func NewRay(origin, dir Vec3) Ray {
	return Ray{origin, dir, 0}
}

func NewRayWithDepth(origin, dir Vec3, depth int) Ray {
	return Ray{origin, dir, depth}
}

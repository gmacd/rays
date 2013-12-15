package renderer

import (
	"gmacd/core"
	"gmacd/geom"
	"gmacd/raytracer"
	"math"
	"runtime"
)

func Render(scene *geom.Scene, canvas *core.Canvas) {
	runtime.GOMAXPROCS(8)

	// init render
	wx1, wx2, wy1, wy2 := -4.0, 4.0, 3.0, -3.0
	sx, sy := 0.0, wy1
	dx := (wx2 - wx1) / float64(canvas.Width)
	dy := (wy2 - wy1) / float64(canvas.Height)

	o := core.NewVec3(0, 0, -5)

	numRays := canvas.Height
	c := make(chan int, numRays)

	for y := 0; y < canvas.Height; y++ {
		sx = wx1

		go func(sx1, sy1 float64, y1 int) {
			for x := 0; x < canvas.Width; x++ {
				dir := core.NewVec3(sx1, sy1, 0).Sub(o).Normal()
				ray := core.NewRay(o, dir)

				raytrace(scene, ray, canvas, x, y1, c)

				sx1 += dx
			}
			c <- 1
		}(sx, sy, y)
		sy += dy
	}

	for i := 0; i < numRays; i++ {
		<-c
	}
}

func raytrace(scene *geom.Scene, ray core.Ray, canvas *core.Canvas, x, y int, c chan int) {
	colour := core.NewColourRGB(0, 0, 0)
	raytracer.Raytrace(scene, ray, &colour, 1, 1.0)
	colour.R = math.Min(colour.R, 1.0)
	colour.G = math.Min(colour.G, 1.0)
	colour.B = math.Min(colour.B, 1.0)
	canvas.SetPixel(x, y, colour)
}

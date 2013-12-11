package renderer

import (
	"gmacd/core"
	"gmacd/geom"
	"gmacd/raytracer"
	"math"
)

func Render(scene *geom.Scene, canvas *core.Canvas) {
	// init render
	wx1, wx2, wy1, wy2 := -4.0, 4.0, 3.0, -3.0
	sx, sy := 0.0, wy1
	dx := (wx2 - wx1) / float64(canvas.Width)
	dy := (wy2 - wy1) / float64(canvas.Height)

	o := core.NewVec3(0, 0, -5)

	for y := 0; y < canvas.Height; y++ {
		sx = wx1

		for x := 0; x < canvas.Width; x++ {
			acc := core.NewColourRGB(0, 0, 0)
			dir := core.NewVec3(sx, sy, 0).Sub(o).Normal()
			ray := core.NewRay(o, dir)

			raytracer.Raytrace(scene, ray, &acc, 1, 1.0)
			acc.R = math.Min(acc.R, 1.0)
			acc.G = math.Min(acc.G, 1.0)
			acc.B = math.Min(acc.B, 1.0)

			canvas.SetPixel(x, y, acc)

			sx += dx
		}
		sy += dy
	}
}

package renderer

import (
	"fmt"
	"gmacd/core"
	"gmacd/geom"
	"gmacd/raytracer"
	"math"
	"runtime"
)

const (
	chunkSizeX int = 16
	chunkSizeY int = 16
)

type Camera struct {
	wx1, wy1 float64
	wx2, wy2 float64
	dx, dy   float64
	origin   core.Vec3
}

func NewCamera(wx1, wy1, wx2, wy2 float64, canvasWidth, canvasHeight int, origin core.Vec3) *Camera {
	dx := (wx2 - wx1) / float64(canvasWidth)
	dy := (wy2 - wy1) / float64(canvasHeight)
	return &Camera{wx1, wy1, wx2, wy2, dx, dy, origin}
}

func (camera *Camera) ScreenToWorld(x, y int) core.Vec3 {
	return core.NewVec3(
		camera.wx1+float64(x)*camera.dx,
		camera.wy1+float64(y)*camera.dy,
		0)
}

func Render(scene *geom.Scene, canvas *core.Canvas) {
	fmt.Printf("Rendering with %v CPUs.\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	camera := NewCamera(-4, 3, 4, -3, canvas.Width, canvas.Height, core.NewVec3(0, 0, -5))

	numXChunks := canvas.Width / chunkSizeX
	numYChunks := canvas.Height / chunkSizeY
	remainderChunkSizeX := canvas.Width - numXChunks*chunkSizeX
	remainderChunkSizeY := canvas.Height - numYChunks*chunkSizeY

	numGoRoutinesSpawned := 0
	c := make(chan int, (numXChunks+1)*(numYChunks+1))

	for chunkY := 0; chunkY < numYChunks; chunkY++ {
		startY := chunkY * chunkSizeY
		endY := startY + chunkSizeY
		for chunkX := 0; chunkX < numXChunks; chunkX++ {
			go renderChunk(
				scene, camera, canvas, c,
				chunkX*chunkSizeX, startY,
				(chunkX+1)*chunkSizeX, endY)
			numGoRoutinesSpawned++
		}
		if remainderChunkSizeX > 0 {
			go renderChunk(
				scene, camera, canvas, c,
				numXChunks*chunkSizeX, startY,
				canvas.Width, endY)
			numGoRoutinesSpawned++
		}
	}
	if remainderChunkSizeY > 0 {
		startY := numYChunks * chunkSizeY
		endY := canvas.Height
		for chunkX := 0; chunkX < numXChunks; chunkX++ {
			go renderChunk(
				scene, camera, canvas, c,
				chunkX*chunkSizeX, startY,
				(chunkX+1)*chunkSizeX, endY)
			numGoRoutinesSpawned++
		}
		if remainderChunkSizeX > 0 {
			go renderChunk(
				scene, camera, canvas, c,
				numXChunks*chunkSizeX, startY,
				canvas.Width, endY)
			numGoRoutinesSpawned++
		}
	}

	for i := 0; i < numGoRoutinesSpawned; i++ {
		<-c
	}
}

func renderChunk(scene *geom.Scene, camera *Camera, canvas *core.Canvas, c chan int, x1, y1, x2, y2 int) {
	for y := y1; y < y2; y++ {
		for x := x1; x < x2; x++ {
			dir := camera.ScreenToWorld(x, y).Sub(camera.origin).Normal()
			ray := core.NewRay(camera.origin, dir)

			raytrace(scene, ray, canvas, x, y, c)
		}
	}
	c <- 1
}

func raytrace(scene *geom.Scene, ray core.Ray, canvas *core.Canvas, x, y int, c chan int) {
	colour := core.NewColourRGB(0, 0, 0)
	raytracer.Raytrace(scene, ray, &colour, 1.0)
	colour.R = math.Min(colour.R, 1.0)
	colour.G = math.Min(colour.G, 1.0)
	colour.B = math.Min(colour.B, 1.0)
	canvas.SetPixel(x, y, colour)
}

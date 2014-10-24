package main

import (
	"fmt"
	"time"

	"github.com/gmacd/rays/core"
	"github.com/gmacd/rays/renderer"
	"github.com/gmacd/rays/testscenes"
)

func main() {
	canvas := core.NewCanvas(640, 480)

	sceneStartTime := time.Now()
	scene := testscenes.CreateTriangleScene()
	sceneEndTime := time.Now()

	renderStartTime := time.Now()
	renderer.Render(scene, canvas)
	renderEndTime := time.Now()

	fmt.Printf("Render complete.\n Scene build took %v.\n Render took %v.\n",
		sceneEndTime.Sub(sceneStartTime),
		renderEndTime.Sub(renderStartTime))

	if err := canvas.Export(); err != nil {
		fmt.Println(err.Error())
	}
}

package main

import (
	"fmt"
	"gmacd/core"
	"gmacd/renderer"
	"gmacd/testscenes"
	"time"
)

func main() {
	canvas := core.NewCanvas(800, 600)

	sceneStartTime := time.Now()
	scene := testscenes.CreateFlipcode3Scene()
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

package main

import (
	"fmt"
	"gmacd/core"
	"gmacd/renderer"
	"gmacd/testscenes"
)

func main() {
	canvas := core.NewCanvas(800, 600)

	scene := testscenes.CreateFlipcode3Scene()

	renderer.Render(scene, canvas)

	if err := canvas.Export(); err != nil {
		fmt.Println(err.Error())
	}
}

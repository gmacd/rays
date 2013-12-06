package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	canvas := NewCanvas(640, 480)

	render(canvas)

	c := color.NRGBA{255, 0, 0, 255}
	draw.Draw(canvas.img, image.Rect(10, 10, 100, 100), &image.Uniform{c}, image.ZP, draw.Src)

	if err := canvas.Export(); err != nil {
		fmt.Println(err.Error())
	}
}

type Canvas struct {
	img    *image.NRGBA
	width  int
	height int
}

func NewCanvas(width, height int) *Canvas {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	canvas := Canvas{img, width, height}
	return &canvas
}

func (canvas *Canvas) SetPixelRGB(x, y int, r, g, b uint8) {
	canvas.img.Set(x, y, color.NRGBA{r, g, b, 255})
}

func (canvas *Canvas) Export() error {
	imgFile, err := os.Create("xxx.png")
	if err != nil {
		return err
	}
	defer imgFile.Close()

	imgWriter := bufio.NewWriter(imgFile)
	if err = png.Encode(imgWriter, canvas.img); err != nil {
		return err
	}
	imgWriter.Flush()

	return nil
}

func render(canvas *Canvas) {
	for y := 0; y < canvas.height; y++ {
		for x := 0; x < canvas.width; x++ {
			canvas.SetPixelRGB(x, y, 0, 0, 0)
		}
	}
}

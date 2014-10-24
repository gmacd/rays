package core

import (
	"bufio"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Canvas struct {
	img    *image.NRGBA
	Width  int
	Height int
}

func NewCanvas(width, height int) *Canvas {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	canvas := Canvas{img, width, height}
	return &canvas
}

func (canvas *Canvas) SetPixelRGB(x, y int, r, g, b uint8) {
	canvas.img.Set(x, y, color.NRGBA{r, g, b, 255})
}

func (canvas *Canvas) SetPixel(x, y int, c ColourRGB) {
	canvas.img.Set(x, y, color.NRGBA{uint8(c.R * 255), uint8(c.G * 255), uint8(c.B * 255), 255})
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

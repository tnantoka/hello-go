package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/StephaneBunel/bresenham"
)

func main() {
	var imgRect = image.Rect(0, 0, 100, 100)
	var img = image.NewRGBA(imgRect)
	var black = color.RGBA{0, 0, 0, 255}

	bresenham.DrawLine(img, 0, 0, 100, 100, black)

	file, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, img)
}

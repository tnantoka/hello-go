package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	length := 10
	pixelRatio := 30
	size := length * pixelRatio

	img := image.NewRGBA(image.Rect(0, 0, size, size))

	colors := []color.RGBA{
		color.RGBA{255, 255, 255, 255},
		color.RGBA{0, 0, 0, 255},
	}

	cells := [][]int{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	for y := range cells {
		for x := range cells[y] {
			for i := 0; i < pixelRatio; i++ {
				for j := 0; j < pixelRatio; j++ {
					img.Set(x*pixelRatio+i, y*pixelRatio+j, colors[cells[y][x]])
				}
			}
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}

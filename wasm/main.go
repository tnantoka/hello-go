package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
)

func main() {
	imageWidth := 100
	imageHeight := 100

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{randomColor()}, image.Point{0, 0}, draw.Src)

	buf := &bytes.Buffer{}
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Stdout.Write(buf.Bytes())
}

func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
}

package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"syscall/js"
)

func main() {
	js.Global().Set("hello", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from", r)
			}
		}()

		size := args[0].Int()
		return generate(size)
	}))

	c := make(chan struct{})
	<-c
}

func generate(size int) string {
	imageWidth := size
	imageHeight := size

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{randomColor()}, image.Point{0, 0}, draw.Src)

	buf := &bytes.Buffer{}
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	encoded := base64.StdEncoding.EncodeToString(buf.Bytes())
	dataURL := fmt.Sprintf("data:image/png;base64,%s", encoded)
	return dataURL
}

func randomColor() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(255)),
		G: uint8(rand.Intn(255)),
		B: uint8(rand.Intn(255)),
		A: 255,
	}
}

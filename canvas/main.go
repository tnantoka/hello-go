package main

import (
	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

func main() {
	c := canvas.New(100, 100)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(canvas.RGBA(255, 255, 255, 1))
	ctx.DrawPath(0, 0, canvas.Rectangle(100, 100))

	ctx.SetStrokeColor(canvas.RGBA(0, 0, 0, 1))
	ctx.DrawPath(0, 0, canvas.Line(100, 100))

	renderers.Write("image.png", c)
}

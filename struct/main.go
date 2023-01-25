package main

import "fmt"

type Point struct {
	X, Y int
}

type Size struct {
	Width, Height int
}

type Rect struct {
	Origin Point
	Size   Size
}

func (self Rect) Description() string {
	return fmt.Sprintf("(Origin: %v, Size: %v)", self.Origin, self.Size)
}

func main() {
	rects := []Rect{
		{Point{0, 0}, Size{100, 100}},
		{Point{10, 10}, Size{200, 200}},
	}

	for i, rect := range rects {
		fmt.Printf("rects[%d] = %s\n", i, rect.Description())
	}
}

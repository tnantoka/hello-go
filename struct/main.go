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

func (self Rect) String() string {
	return self.Description()
}

func main() {
	rects := []Rect{
		{Point{0, 0}, Size{100, 100}},
		{Point{10, 10}, Size{200, 200}},
	}

	for i, rect := range rects {
		fmt.Printf("rects[%d]: ", i)
		fmt.Println(rect)
	}

	rect1 := Rect{Point{0, 0}, Size{100, 100}}
	rect2 := rect1
	rect1.Origin.X = 10
	fmt.Printf("1: %v, 2: %v\n", rect1.Origin.X, rect2.Origin.X)

	rect3 := &Rect{Point{0, 0}, Size{100, 100}}
	rect4 := rect3
	rect3.Origin.X = 10
	fmt.Printf("3: %v, 4: %v\n", rect3.Origin.X, rect4.Origin.X)

}

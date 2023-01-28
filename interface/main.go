package main

import (
	"fmt"
)

type ExampleInterface interface {
	Example()
}

type ExampleStruct struct {
}

func (e ExampleStruct) Example() {
}

func ExampleFunc1(i int64) ExampleInterface {
	if i == 0 {
		return ExampleFunc2()
	} else {
		return ExampleFunc3()
	}
}

func ExampleFunc2() ExampleStruct {
	return ExampleStruct{}
}
func ExampleFunc3() *ExampleStruct {
	return &ExampleStruct{}
}

func main() {
	fmt.Println(ExampleFunc1(0))
	fmt.Println(ExampleFunc1(1))

	var example1 ExampleInterface = ExampleStruct{}
	var example2 ExampleInterface = &ExampleStruct{}
	fmt.Println(example1)
	fmt.Println(example2)
}

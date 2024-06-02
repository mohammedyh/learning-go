package main

import "fmt"

type Triangle struct {
	height float32
	base   float32
}

func (t Triangle) area() float32 {
	return 0.5 * (t.base * t.height)
}

func main() {
	triangle := Triangle{10, 4}

	fmt.Println(triangle)
	fmt.Println(triangle.area())
}

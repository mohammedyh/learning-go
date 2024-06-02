package main

import "fmt"

type Triangle struct {
	height float32
	base   float32
}

func (t *Triangle) updateBase(newBase float32) {
	t.base = newBase
}

func main() {
	triangle := Triangle{10, 4}
	triangle.updateBase(13)

	fmt.Println(triangle)
}

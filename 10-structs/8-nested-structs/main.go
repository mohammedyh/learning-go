package main

import "fmt"

type Point struct {
	x int
	y int
}

type Circle struct {
	// point Point
	Point
	radius int
}

func main() {
	circle := Circle{
		Point{4, 5},
		2,
	}

	// fmt.Println(circle.point.x)
	fmt.Println(circle.x)
}

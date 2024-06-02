package main

import "fmt"

type country struct {
	name      string
	capital   string
	latitude  float32
	longitude float32
}

func main() {
	france := country{}
	fmt.Println(france)
}

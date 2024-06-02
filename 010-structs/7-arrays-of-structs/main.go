package main

import "fmt"

type Cake struct {
	typeOfCake string
	weight     int
}

func main() {
	cakes := []Cake{{"Chocolate", 1000}, {"Carrot", 500}, {"Ice Cream", 2000}}

	cakes[1].weight = 1500

	fmt.Println(cakes[0].weight)
	fmt.Println(cakes)
}

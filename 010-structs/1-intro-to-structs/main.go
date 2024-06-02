package main

import "fmt"

// Example of a struct definition
type Pet struct {
	name    string
	petType string
	age     int
}

func main() {
	// Using a struct type
	nuggets := Pet{"Nuggets", "dog", 4}
	muffin := Pet{"Muffin", "cat", 7}
	robin := Pet{"Robin", "bird", 2}

	fmt.Println(nuggets)
	fmt.Println(muffin)
	fmt.Println(robin)
}

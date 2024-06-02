package main

import "fmt"

func main() {
	star := "Polaris"
	// Declared with var and type for clarity
	var starPointer *string = &star

	fmt.Println("The address of star is", starPointer)
}

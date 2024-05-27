package main

import "fmt"

func main() {
	star := "Polaris"
	starAddress := &star

	*starAddress = "Sirius"

	fmt.Println("The actual value of star is", star)
}

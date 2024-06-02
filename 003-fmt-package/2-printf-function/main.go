package main

import "fmt"

func main() {
	animal1, animal2 := "Cat", "Dog"

	fmt.Printf("Are you a %v or a %v person?\n", animal1, animal2)
	// Output: Are you a Cat or a Dog person

	// The equivalent using fmt.Println
	fmt.Println("Are you a", animal1, "or a", animal2, "person?")
}

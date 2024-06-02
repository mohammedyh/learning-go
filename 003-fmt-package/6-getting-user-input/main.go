package main

import "fmt"

func main() {
	// Single argument / value
	fmt.Println("How are you doing?")

	var response string
	fmt.Scan(&response)

	fmt.Printf("I'm %v.", response)
	// Output: How are you doing?
	// Input: good
	// Output: I'm good

	// Multiple arguments / values
	fmt.Println("How are you doing?")

	var response1, response2 string
	fmt.Scan(&response1)
	fmt.Scan(&response2)

	// Can also be written more concisely
	// fmt.Scan(&response1, &response2)

	fmt.Printf("I'm %v %v, thanks for asking\n", response1, response2)
	// Output: How are you doing?
	// Input: very well
	// Output: I'm very well, thanks for asking
}

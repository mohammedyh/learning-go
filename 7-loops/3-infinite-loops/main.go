package main

import "fmt"

// This is outside count because it would otherwise keep
// getting initialized to 0 each time count is called
var number int

func count() {
	var input int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)

	number = number + input
	fmt.Println("Total guests:", number)
}

func main() {
	for {
		count()
	}
}

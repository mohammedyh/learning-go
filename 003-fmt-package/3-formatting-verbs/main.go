package main

import "fmt"

func main() {
	floatExample := 1.75
	fmt.Printf("Working with a %T", floatExample)
	// Output: Working with a float64

	yearsOfExp := 3
	reqYearsExp := 15
	fmt.Printf("I have %d years of Go experience and this job is asking for %d years", yearsOfExp, reqYearsExp)
	// Output: I have 3 years of Go experience and this job is asking for 15 years

	stockPrice := 3.5423
	fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)
	// Output: Each share of Gopher feed is 3.54!

	fmt.Printf("Working with a ____", floatExample)
	// Output: Working with a ____%!(EXTRA float64=1.75).

}

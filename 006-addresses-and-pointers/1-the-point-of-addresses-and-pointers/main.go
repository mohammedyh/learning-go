package main

import "fmt"

func addHundred(num int) {
	num += 100
}

func main() {
	x := 1

	addHundred(x)

	fmt.Println(x)
	// Output: 1
	// But why?
}

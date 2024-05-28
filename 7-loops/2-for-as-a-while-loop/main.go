package main

import "fmt"

func ask() int {
	var input int
	fmt.Print("I am thinking of a number between 1-100: ")
	fmt.Scan(&input)
	return input
}

func main() {
	var guess int

	// Continually ask until input is 56
	for guess != 56 {
		guess = ask()
	}

	if guess == 56 {
		fmt.Println("You are correct! You may go through to the treasure!")
	}
}

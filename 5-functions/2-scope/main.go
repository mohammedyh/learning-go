package main

import "fmt"

func startGame() {
	instructions := "Press enter to start..."

	fmt.Println(instructions)
}

func main() {
	startGame()
	// fmt.Println(instructions)
	// Output: undefined: instructions
}

package main

import "fmt"

func main() {
	// Variable declaration
	var lengthOfSong uint16 // 0 - 65,536
	var isMusicOver bool    // 0 or 1
	var songRating float32  // 32 bit precision

	// Variable assignment
	lengthOfSong = 10000
	isMusicOver = true
	songRating = 3.72

	// Variable initialization (declaration & assignment in one)
	var kilometersToMars int32 = 62100000

	fmt.Println(lengthOfSong)
	fmt.Println(isMusicOver)
	fmt.Println(songRating)
	fmt.Println(kilometersToMars)

	// ------------
	// Assignment operators

	var basketTotal float64
	carrotPrice := 0.75

	basketTotal += carrotPrice
	fmt.Println(basketTotal) // Prints: 0.75

	// ------------
	// Multiple variable declaration

	// Using this syntax, both variables must be the same type
	var part1, part2 string
	part1 = "To be..."
	part2 = "Not to be..."
	fmt.Println(part1, part2)

	quote, fact := "Bears, Beets, Battlestar Galactica", true
	fmt.Println(quote) // Prints: Bears, Beets, Battlestar Galactica
	fmt.Println(fact)  // Prints: true

}

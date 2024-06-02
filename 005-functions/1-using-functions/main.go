package main

import "fmt"

func summonAI() {
	fmt.Println("Hey there, I'm ChatGPT, how can I help?")
}

func main() {
	summonAI()
	summonAI()

	addNums := func(x, y int) int {
		return x + y
	}

	fmt.Println("result:", addNums(2, 2))
}

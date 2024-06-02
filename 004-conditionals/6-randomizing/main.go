package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	// Set the below env variable to disable
	// autoseeding introduced as of Go 1.20
	os.Setenv("GODEBUG", "randautoseed=0")

	amountLeft := rand.Intn(10000)
	amountLeft2 := rand.Intn(10000)

	fmt.Println("amountLeft is: ", amountLeft)
	fmt.Println("amountLeft2 is: ", amountLeft2)

	if amountLeft > 5000 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}

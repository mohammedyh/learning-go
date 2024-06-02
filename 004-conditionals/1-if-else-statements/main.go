package main

import "fmt"

func main() {
	heistReady := true

	if heistReady {
		fmt.Println("Let's Go!")
	} else {
		fmt.Println("Act normal")
	}

	amountStolen := 64650

	if amountStolen > 1000000 {
		fmt.Println("We've hit the jackpot!")
	} else if amountStolen >= 5000 {
		fmt.Println("Think of all the candy we can buy!")
	} else {
		fmt.Println("Why did we even do this?")
	}
}

package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	fmt.Println(donuts)

	firstChoice := donuts["frosted"]
	secondChoice, status := donuts["bavarian cream"]

	fmt.Println(firstChoice)
	fmt.Println(secondChoice)

	if status {
		fmt.Println(len(donuts))
	} else {
		fmt.Println("We do not sell that donut!")
	}
}

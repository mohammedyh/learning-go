package main

import "fmt"

func main() {
	menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

	fmt.Println("The menu:")
	for index, item := range menu {
		fmt.Printf("Item: %v: %v\n", index+1, item)
	}

	orders := map[string]string{
		"John":   "Cheeseburgers",
		"Janet":  "Hamburgers",
		"Jordan": "Fries",
	}
	orders["James"] = "Chicken Sandwich"

	fmt.Println("\nThe friend's orders:")

	for person, order := range orders {
		fmt.Printf("%v ordered %v\n", person, order)
	}
}

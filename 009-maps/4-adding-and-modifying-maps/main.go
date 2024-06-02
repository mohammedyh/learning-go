package main

import "fmt"

func main() {
	donuts := map[string]int{
		"frosted":   10,
		"chocolate": 15,
		"jelly":     8,
	}

	fmt.Println(donuts)

	donuts["glazed"] = 12
	donuts["jelly"] = 3

	fmt.Println(donuts["glazed"])
	fmt.Println(donuts)
}

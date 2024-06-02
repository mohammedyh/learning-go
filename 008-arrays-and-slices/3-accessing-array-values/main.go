package main

import "fmt"

func main() {
	triangleAngles := [3]int{30, 60, 90}
	var sum int

	// Use a loop to calculate the sum, we can do better than
	// manually access the index and adding them together :)
	// sum = triangleAngles[0] + triangleAngles[1] + triangleAngles[2]

	for _, val := range triangleAngles {
		sum += val
	}

	fmt.Println("triangleAngles[2]:", triangleAngles[2])
	fmt.Println("sum:", sum)
}

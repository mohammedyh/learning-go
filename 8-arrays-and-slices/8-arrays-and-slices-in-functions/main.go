package main

import "fmt"

func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}
	changeLastElement(myTutors, "Bobby")
}

func changeLastElement(slice []string, tutor string) {
	if len(slice) > 0 {
		slice[len(slice) - 1] = tutor
	}
	fmt.Println(slice)
}

package main

import "fmt"

func main() {
	myTutors := []string{"Kirsty", "Mishell", "Jose", "Neil"}

	fmt.Println("Length:", len(myTutors))
	fmt.Println("Capacity:", cap(myTutors))
}

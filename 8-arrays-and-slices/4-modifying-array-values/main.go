package main

import "fmt"

func main() {
	myDogs := [3]string{"Frida", "Fedo", "Jegf"}
	myDogs[1] = "Fredo"
	myDogs[2] = "Jeff"

	fmt.Println(myDogs)
}

package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
	grade     int
}

func main() {
	peter := Student{"Peter", "Bookman", 16, 11}
	scott := Student{firstName: "Scott", lastName: "Peterson", grade: 12}

	fmt.Println(peter)
	fmt.Println(scott)
}

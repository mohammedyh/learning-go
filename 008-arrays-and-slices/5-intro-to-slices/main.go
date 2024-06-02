package main

import "fmt"

func main() {
	myTutors := [4]string{"Kirsty", "Mishell", "Jose", "Neil"}
	tutorsSlice := myTutors[:]
	tutorSubjects := []string{"Go", "Java", "Python"}

	fmt.Println("tutorsSlice:", tutorsSlice)
	fmt.Println("tutorSubjects:", tutorSubjects)
	fmt.Println("myTutors:", myTutors)

	// The slice elements point to the same memory addresses
	// as the elements in the original array
	fmt.Println("myTutors:", &myTutors[0], &myTutors[1], &myTutors[2], &myTutors[3])
	fmt.Println("tutorsSlice:", &tutorsSlice[0], &tutorsSlice[1], &tutorsSlice[2], &tutorsSlice[3])
}

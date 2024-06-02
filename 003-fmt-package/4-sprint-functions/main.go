package main

import "fmt"

func main() {
	grade := "100"
	compliment := "Great job!"
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

	fmt.Println(teacherSays)
	// Output: You scored a 100 on the test! Great job!
}

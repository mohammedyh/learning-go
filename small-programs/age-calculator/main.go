package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Enter your birth year: ")

	var userYear string
	const PURPLE = "\033[35m"
	const NOCOLOR = "\033[0m"

	fmt.Scan(&userYear)

	userYearInt, err := strconv.ParseInt(userYear, 10, 16)
	currentYear := time.Now().Year()

	if err != nil {
		log.Fatalf("%v%v%v", PURPLE, err.Error(), NOCOLOR)
	}

	if int(userYearInt) > currentYear {
		log.Fatalf("%vThe birth year cannot be in the future%v", PURPLE, NOCOLOR)
	}

	age := currentYear - int(userYearInt)

	fmt.Printf("%vYou will be %v years old this year%v\n", PURPLE, age, NOCOLOR)
}

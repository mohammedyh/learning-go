package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Printf("Enter your birth year: ")

	var year string

	fmt.Scan(&year)

	yearAsInt, err := strconv.ParseInt(year, 10, 16)

	if err != nil {
		log.Fatal(err.Error())
	}

	if int(yearAsInt) > time.Now().Year() {
		log.Fatal("The birth year cannot be in the future")
	}

	age := time.Now().Year() - int(yearAsInt)

	fmt.Printf("You will be %v years old this year\n", age)
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// using short declaration operator
	daysOnVacation := 8
	// alternative syntax using var
	var hoursInDay = 24

	fmt.Println("You have spent", daysOnVacation*hoursInDay, "hours on vacation.")

	// Will print the architecture like arm (32 is implied), arm64 etc.
	fmt.Println(runtime.GOARCH)
}

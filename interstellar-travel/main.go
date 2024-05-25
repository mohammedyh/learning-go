package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func fuelGauge(fuel int) {
	// Prints number formatted with commas
	p := message.NewPrinter(language.English)
	p.Printf("You have %v gallons of fuel remaining.\n", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300_000
	case "Mercury":
		fuel = 500_000
	case "Mars":
		fuel = 700_000
	default:
		fuel = 0
	}

	return fuel
}

func greetPlanet(planet string) {
	// Unicode code point for planet emoji
	// Copied as: U+1FA90
	// Replace + with 000
	// and prefix with \ (as a rune)
	planetEmoji := '\U0001FA90'
	fmt.Printf("Your destination is %v %c We hope you have a nice visit\n", planet, planetEmoji)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	fuelRemaining := fuel
	fuelCost := calculateFuel(planet)

	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	fuel := 1_000_000
	planetChoice := "Mercury"

	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}

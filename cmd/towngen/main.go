package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/ironarachne/towngen"
)

func displayTown(town towngen.Town) {
	fmt.Println(town.Name)
	fmt.Println("Category: " + town.Category.Name)
	fmt.Println("Population: " + strconv.Itoa(town.Population))
	fmt.Println("Mayor: " + town.Mayor.FirstName + " " + town.Mayor.LastName)
	fmt.Println("Exports")
	fmt.Println("-------")

	for good, quantity := range town.Exports {
		fmt.Println("- " + good + " (Quantity: " + strconv.Itoa(quantity) + ")")
	}

	fmt.Println("Imports")
	fmt.Println("-------")

	for good, quantity := range town.Imports {
		fmt.Println("- " + good + " (Quantity: " + strconv.Itoa(quantity) + ")")
	}
}

func main() {
	randomSeed := flag.Int64("s", 0, "Optional random generator seed")

	flag.Parse()

	if *randomSeed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(*randomSeed)
	}

	town := towngen.GenerateTown()

	displayTown(town)
}
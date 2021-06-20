package main

import (
	"fmt"
)

func main() {
	// "IF" statement
	var fruit string

	fmt.Println("Enter the fruit name, It will print the month name\n 1. Orange\n 2. Mango\n 3. Banana\n 4. Papaya\n 5. Guava\n")
	fmt.Scanln(&fruit)
	if fruit == "Orange" {
		fmt.Println("Month of September")
	} else if fruit == "Mango" {
		fmt.Println("Month of June")
	} else if fruit == "Banana" {
		fmt.Println("Month of March")
	} else if fruit == "Papaya" {
		fmt.Println("Month of August")
	} else if fruit == "Guava" {
		fmt.Println("Month of October")
	} else {
		fmt.Println("No fruits from the list")
	}
}


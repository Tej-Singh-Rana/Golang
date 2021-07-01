package main

import "fmt"

func main() {
	// main code block

	book := "Himalayan Travels"
	var prize int
	fmt.Println("Enter the amount to purchase the book :- ")
	fmt.Scanln(&prize)

	if prize < 300 {
		fmt.Println("Added into the cart book:", book)
	} else {
		fmt.Printf("Amount is not sufficient to purchase %s.", book)
	}
}

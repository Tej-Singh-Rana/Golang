package main

import (
	"fmt"
)

func main() {
	// By using the short variable declaration operator (foo := "bar" --> string, foo := 5 --> int)
	firstName := "Roni"
	lastName := "Walker"
	roomNumber := 45

	fmt.Println("Name is", firstName, lastName, "and room number is", roomNumber, ".")
}

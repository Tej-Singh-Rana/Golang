package main

import (
	"fmt"
)

// Declare and Initialize variables before main func

var (
	firstName, lastName string = "John", "Wick"
	//lastName  string = "Wick"
)

func main() {
	fmt.Println("Introducing stories of", firstName, lastName,".")
}

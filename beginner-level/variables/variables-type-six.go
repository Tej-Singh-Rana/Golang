package main

import (
	"fmt"
)

// Variables will take a type of the initialized values. e.g. string (""), int (number)
var (
	developer = "Self"
	language  = "Go"
	year      = 2020
)

func main() {
	// main func block where we call it first and last.
	fmt.Printf("%s learning %s language since %d.", developer, language, year)

}

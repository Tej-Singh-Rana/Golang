package main

import (
	"fmt"
)

func main() {
	// main func block
	var goose = "duck duck go!!"   // Declared variable without initialized it's type (like string, int, float).
	var limit = "Unlimited ride!!" // Go ecosystem will automatically detect the type of these variables.
	var speed = 150
	fmt.Println(goose + " " + limit)
	fmt.Println("Speed of goose", speed)
}

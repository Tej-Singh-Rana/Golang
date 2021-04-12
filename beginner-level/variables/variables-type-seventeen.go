package main

import (
	"fmt"
)

// Declare function "math" with "func" keyword
func math(x, y int) int {
	return x + y // return the sum of "x" and "y" integer value.
}

func main() {
	total := math(-2, 10) // Assign the "math" function to the variable.

  fmt.Println("Total value of x, y = ", total) // It will print the sum value (x+y).
}

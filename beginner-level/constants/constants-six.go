package main

import (
	"fmt"
)

func main() {
	// Declare and initialize constant var "i" and "y" in a single line
	const i, y = 4, "word"
	fmt.Println(i, y)

	// cannot assign to i (declared const)
	// i = 12
	// fmt.Println(i)
}

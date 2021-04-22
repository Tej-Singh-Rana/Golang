package main

import (
	"fmt"
)

func main() {

	// Slice + For Loop

	cold := []int{1, 2, 3, 4, 5}

	for index, value := range cold {

		fmt.Println("Print the index = ", index, "| Print the element = ", value)
	}
}


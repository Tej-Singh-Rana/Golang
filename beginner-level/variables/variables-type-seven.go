package main

import (
	"fmt"
	"strconv"
)

var core int
var studio string

func main() {
	// Specify type and perform a conversion
	// Conversion int to "string"
	core = 12
	core1 := strconv.Itoa(core)
	fmt.Printf("core1 value = %s and core1 type = %T", core1, core1)
	fmt.Println("")

	// Conversion "string" to int
	studio = "50"
	if studioCart, err := strconv.Atoi(studio); err == nil {
		// %v -> print the value in a default format, %T -> print the type of the value.
		fmt.Printf("studioCart value = %v and studioCart type = %T", studioCart, studioCart)
	}
}

package main

import (
	"fmt"
)

func main() {
	// Main code block
	// Maps in Go
	// map[KeyType]ValueType
	
	var x map[string]int
	
	// Not initialized any value to "x"
	// zero value of a map is nil.
	fmt.Printf("Type %T and value %v", x, x)
}

/*
_Output_:-
Type map[string]int and value map[]

*/

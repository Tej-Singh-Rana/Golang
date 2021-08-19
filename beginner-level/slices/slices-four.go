package main

import (
	"fmt"
)

func main() {
	// Main code block

	// Every element in the slices should be of the same type.
	// In below example, declared "int" value. Because of that we are getting an error.
	// Error :- cannot use 2403 (type untyped int) as type string in slice literal

	// var x []string = []string{"world", "famous", "place", 2403, "is", "jaipur"}

	// Trying after removing the integer.

	var x []string = []string{"World", "famous", "city", "is", "Jaipur."}
	for _, i := range x {

		fmt.Println(i)

	}

}

/*
_Output_:-
World
famous
city
is
Jaipur.

*/

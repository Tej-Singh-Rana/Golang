package main

import (
	"fmt"
)

// empty interface

type First interface{}

func main() {
	// Main code block

	var t First

	fmt.Println("Value of empty interface", t)

}

/*
_Output_:-
Value of empty interface <nil>


*/

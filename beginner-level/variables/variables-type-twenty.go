package main

import (
	"fmt"
)

// Declared multiple values
var (
	exprOne string
	exprTwo int
	status  bool
)

func main() {
	// Main code block
	//Initialized values
	exprOne = "John"
	exprTwo = 2021
	status = false

	fmt.Printf("%s details on %d and it's condition %t.", exprOne, exprTwo, status)

	// Changed the type
	var (
		exprOne int
		exprTwo string
		status  bool
	)

	//Initialized values with different type.
	exprOne = 2021
	exprTwo = "Rocky"
	status = true

	fmt.Printf("\n%s details on %v and it's condition %t.", exprTwo, exprOne, status)

}


/*
_Output_:-
John details on 2021 and it's condition false.
Rocky details on 2021 and it's condition true.

*/

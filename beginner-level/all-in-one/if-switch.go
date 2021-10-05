package main

import (
	"fmt"
)

func main() {
	// Main code block
	var x string = "Yes"
	x = "No"

	if x == "Yes" || x == "No" {

		switch x {

		default:
			fmt.Println("Find out the correct!")

		case "Yes":
			fmt.Println("You are correct!")
		}

	}
}

/*
_Output_:-
Find out the correct!


*/

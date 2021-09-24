// https://play.golang.org/p/IkI7_Pbya-4

package main

import (
	"fmt"
)

func main() {
	// Main code block

	x := 0

	if x < 15 {
		fmt.Println("lt 15")
	} else {
		// goto statement

		goto Cell

	}

	for x < 12 {
		fmt.Println(x)
		x++

		// Label is declared in the for loop block and we are specifying the goto statement in the above if/else statement.
		// So because of curly bracket blocks it is not possible to jump over it.
	Cell:
		fmt.Println("Value should be less than 15")

	}
}

/*
_Output_:-
line:19: goto Cell jumps into block {curly brackets} starting at line:23

*/

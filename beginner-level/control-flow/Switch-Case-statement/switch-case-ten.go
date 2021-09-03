// Go playground - https://play.golang.org/p/86BEjA4QilJ
package main

import (
	"fmt"
)

// This one will be evaluated first.
func Vision(x int) int { fmt.Println(x); fmt.Println("Looping"); return x }

func main() {
	// Main code block

	switch Vision(10) {

	// Evaluated left to right and top to bottom.
	// Each value will be gone through the Vision() func first. As you can see from the Output result.
	case Vision(5), Vision(10), Vision(15):
		fmt.Println("Final End") // --> At the end, it will print "Final End".

	case Vision(20):
		fmt.Println("Final Chapter")
	}

}

/*
_Output_:-
10
Looping
5
Looping
10
Looping
Final End

*/

package main

import (
	"fmt"
)

func main() {
	// Main code block

	// Anonymous Function
	// We can assign it to the variable
	number := func() {
		var x int = 12345
		fmt.Println(x)
	}
	
	// Now, you can call this variable as a func
	number()
}


/*
_Output_:-
12345


*/

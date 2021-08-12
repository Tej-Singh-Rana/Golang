package main

import (
	"fmt"
)

func main() {
	// Main code block
       // type *int is a pointer to a "x" value
	var x *int

	y := 12
	// Address the value
	x = &y
    
	fmt.Println(*x)

}

/*
_Output_:-
12

*/

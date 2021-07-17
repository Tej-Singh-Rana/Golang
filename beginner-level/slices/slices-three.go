package main

import (
	"fmt"
)

func main() {
   // Main code block
  // Slice literal, where we don't specify length inside square brackets[]. 

	var student = []string{
		"john",
		"jacob",
		"rahul",
		"raghu",
		"narayan",
	}

	// for loop with range function
	for _, count := range student {
		fmt.Println(count)
	}
}


/*
_Output_:-
john
jacob
rahul
raghu
narayan

*/

package main

import (
	"fmt"
)

func main() {
	// Main code block
        
	gradeA := "pass"
	
	// Declares pointer 
	var gradeB *string

	gradeB = &gradeA
	fmt.Println(&gradeA, &gradeB)

}


/*
_Output_:-
0xc00009e210 0xc0000b2018

*/

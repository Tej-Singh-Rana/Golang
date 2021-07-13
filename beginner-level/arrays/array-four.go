package main

import (
	"fmt"
)

func main() {
	// Main code block
	// Name := [Length]Type{value-1, value-2, ..., N}
	
	// short declaration
	language := [4]string{"go", "javascript", "python", "R"}

	fmt.Println(language)
 	fmt.Println(language[0])
	fmt.Println(language[3])
  
}

/*
_Output_:-
[go javascript python R]
go
R

*/


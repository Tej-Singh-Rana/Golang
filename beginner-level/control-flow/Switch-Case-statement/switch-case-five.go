package main

import (
	"fmt"
)

var (
	factor int = 5
)

func main() {
	// Main code block
	// If value matched then it won't move to the next line, even that's resulting to true.
	switch factor {

	// Assign it multiple values to match.
	case 1, 2, 3:
		fmt.Println("1, 2, 3")

	case 4, 5:
		fmt.Println("4 and 5")

	// You can only have one case statement for each value.
	// If we will uncomment this, It will give an error.
	// case 5:
	//	fmt.Println("Perfect!!")

	default:
		fmt.Println("Recheck the number.")

	}

}

/*
_Output_:-
4 and 5

*/

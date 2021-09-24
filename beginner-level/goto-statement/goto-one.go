// https://play.golang.org/p/Cb-cqEWJGPY

package main

import (
	"fmt"
)

func main() {
	// Main code block

	number := 1
	nameMatch := "Rock"

	if number == 1 {
		// Jump into the label called "newCall" and exit out after that.

		goto newCall

		// It will not print the below message
		// fmt.Println("It will Print?")

	} else {
		fmt.Println("Try again!!")
	}

newCall:
	if nameMatch == "Rocky" {
		fmt.Println("name matched.")
	} else {
		fmt.Println("Not found in the record.")
	}

}

/*
_Output_:-
Not found in the record.

*/


package main // declared main package

// imported fmt package
import (
	"fmt"
)

var examYear int = 2017

func main() {
	// main func body

	// "default:" case used when no case matches.
	switch {
	case examYear == 2018:
		fmt.Println("Exam year was 2018.")
	case examYear == 2020:
		fmt.Println("Exam year was 2020.")

		// No matches found so default should be printed.
	default:
		fmt.Println("None of the year match.")
	}

}

/*
_Output_:-
None of the year match.

*/

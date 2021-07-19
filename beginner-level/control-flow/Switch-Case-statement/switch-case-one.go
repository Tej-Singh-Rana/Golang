package main // declared main package

// imported fmt package
import (
	"fmt"
)

func main() {
	// main func body
	// Simple switch case statement
	var name = "John"
	// changed the value
	name = "Ram"
	// So var "name" value will match to each case. If it's true then it will not move to the next line, even if second line is true.
	switch name {
	case "Ram":
		fmt.Println("This is Ram.")
	case "John":
		fmt.Println("This is John.")
	}

}

/*
Output :-
This is Ram.

*/

package main // declared main package

// imported fmt package
import (
	"fmt"
)

func main() {
	// main func body
	// Simple switch case statement
	var name = "John"
	name = "Ram"
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

package main

import "fmt"

func main() {
	// code block
	goo := 10
	switch goo {
	case 10:
		fmt.Println("Correct case.")
	case 20:
		fmt.Println("Incorrect case.")
	default:
		fmt.Println("None")
	}
}

/*
Output :-

Correct case.
*/

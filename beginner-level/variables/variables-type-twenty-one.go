package main

import "fmt"

// Declared and initialized func to variable

var color = func() string {
	return "red"
}

func main() {
	// Main code block
	fmt.Println(color())

}

/*
_Output_:-
red

*/

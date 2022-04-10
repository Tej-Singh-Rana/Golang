package main

import "fmt"

func text() string {
	x := "Hello! Please accept my request."
	return x
}
func main() {
	// Main code block

	request := text()
	if request == "Hello! Please accept my request." {
		fmt.Println("Message matched.")
	} else {
		fmt.Println("Message not matched.")
	}
}

/*
_Output_:-
Message matched.

*/

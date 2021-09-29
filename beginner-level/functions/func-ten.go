// https://play.golang.org/p/HShQqpFVWKJ

package main

import "fmt"

func main() {
	// Main code block

	CodeBlock("This is a testing message!!")

}

// Print message

func CodeBlock(p string) string {

	fmt.Println("A message you wrote:", p)

	return p

}

/*
_Output_:-
A message you wrote: This is a testing message!!

*/

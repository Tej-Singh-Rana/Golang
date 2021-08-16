package main

import (
	"fmt"
)

// By default, we have two function main() and init().
// It's same as main() but it executes first and we can define multiple times as we want.
// function init() must have no arguments and no return values.
// func init(x int) int { return } --> Not supported.

func init() {
	// Init code block
	fmt.Println("First init block")

}

func main() {
	// Main code block
	fmt.Println("Main block")
}

// It follows the ordering.

func init() {
	// Init code block
	fmt.Println("Third init block")

}

func init() {
	// Init code block
	fmt.Println("Second init block")

}

func init() {
	// Init code block
	fmt.Println("Fourth init block")

}



/*
_Output_:-
First init block
Third init block
Second init block
Fourth init block
Main block

*/

// https://play.golang.org/p/heUxweIrbQq

package main

import (
	"fmt"
)

// Declare and initialize, type map[string]string
var tempClass map[string]string

func main() {
	// Main code block
	// maps with make function
	tempClass = make(map[string]string)

	tempClass["first"] = "High"
	tempClass["second"] = "Medium"
	tempClass["third"] = "Average"
	tempClass["four"] = "Low"

	// Type and Value
	fmt.Printf("Type of variable = %T, key-value pairs = %v\n", tempClass, tempClass)

	// underScore(_) will ignore that index value and will not be used in the std output.
	for _, value := range tempClass {
		fmt.Println("tempClass values = ", value)
	}
}

/*
_Output_:-
Type of variable = map[string]string, key-value pairs = map[first:High four:Low second:Medium third:Average]
tempClass values =  High
tempClass values =  Medium
tempClass values =  Average
tempClass values =  Low

*/

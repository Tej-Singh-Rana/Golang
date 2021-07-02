package main

import (
	"fmt"
)

// Declared new Type

type Number int

type Role string

func main() {
	var name Role = "SRE"
	var position Number = 1

	fmt.Println(name, position)
	fmt.Printf("Type = %T, %T", name, position)
}

/*
_Output_:-
SRE 1
Type = main.Role, main.Number
*/

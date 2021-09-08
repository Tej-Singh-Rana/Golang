package main

import "fmt"

// type struct
type foo struct {
	Name   string
	Class  string
	Number int64
}

func main() {
	// Main function body

	call := foo{
		Name:   "Rahul",
		Class:  "First",
		Number: 12,
	}

	fmt.Println(call)
	// Printing individual values
	fmt.Println(call.Name)
	fmt.Println(call.Class)
	fmt.Println(call.Number)

}

/*
_Output_:-
{Rahul First 12}
Rahul
First
12

*/

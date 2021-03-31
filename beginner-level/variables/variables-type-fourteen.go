package main

import "fmt"

// type struct
type card struct {
	ID         int
	Name       string
	RollNumber int
	pass       bool
}

// Declare variable "result" and initialize slice "card" to it.
var result = []card{
	{ID: 1, Name: "Josh", RollNumber: 1200, pass: true},
	{ID: 2, Name: "Tej", RollNumber: 1201, pass: false},
	{ID: 3, Name: "John", RollNumber: 1202, pass: true},
}

func main() {
	fmt.Println(result)
}



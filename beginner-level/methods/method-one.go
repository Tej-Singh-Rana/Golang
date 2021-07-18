package main

import (
	"fmt"
)

// Declared own struct type
type Bus struct {
	route string
	name  string
}

// Methods
func (b Bus) info() string {
	return "Name of the bus is " + b.name + " and route of that bus is " + b.route +"."
}

func main() {
	raw := Bus{name: "Plaza", route: "16"}
	fmt.Println(raw.info())
}

/*
_Output_:-
Name of the bus is Plaza and route of that bus is 16.

*/

package main

import (
	"fmt"
)

// variable declared globally

var tool int = 20

func main() {
  // change the globally declared value
	tool = 30
  
  // declared value for "atom" variable
	if atom, err := fmt.Printf("Value of 'tool' = %v", tool); err == nil {
		fmt.Printf("%v", atom)
	}
}

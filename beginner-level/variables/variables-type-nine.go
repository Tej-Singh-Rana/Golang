package main

import "fmt"

func main() {
	// Variable can contain functions.

	test := func() {
		// function body
		// write code
		fmt.Println("Write and store the value in a variable 'test'.")
	}
	test()
	
	fmt.Printf("Type of variable 'test' = %T", test)
}

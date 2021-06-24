package main

import (
	"fmt"
)

// Main entry point

func main() {
	// Main func body
	var token int
	fmt.Scanln(&token)
	var name string
	fmt.Scanln(&name)

	if token == 1 && name == "ram" {
		token = 1 * 5
		fmt.Printf("Token number is %v", token)
	} else if token == 2 && name == "ali" {
		token = 2 * 5
		fmt.Printf("Token number is %v", token)
	} else if token == 3 && name == "john" {
		token = 3 * 5
		fmt.Printf("Token number is %v", token)
	} else if token == 4 && name == "hari" {
		token = 4 * 5
		fmt.Printf("Token number is %v", token)
	} else if token == 5 && name == "tej" {
		token = 5 * 5
		fmt.Printf("Token number is %v", token)
	} else {
		fmt.Println("Not found.")
	}

}

/*
Output :-

Token number is 25

*/

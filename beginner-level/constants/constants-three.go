package main

import "fmt"

// Declare multiple constant variables
const (
	roleOne   string = "Dev"
	roleTwo   string = "Admin"
	roleThree string = "Ops"
)

func main() {
	// Declare inside main func
	fmt.Println("Roles", "\n", roleOne, "\n", roleTwo, "\n", roleThree)

}

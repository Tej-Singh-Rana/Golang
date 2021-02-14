package main

import (
	"fmt"
)

// Declared variables separately

var department string
var firstName string
var lastName string
var percentage float64
var pass bool

func main() {
	department = "Math"
	firstName = "Ravi"
	lastName = "Ram"
	percentage = 74.5589898
	pass = true
	fmt.Printf("Department\tFirst Name\tLast Name\tPercentage\tPass\t\t\n%s\t\t%s\t\t%s\t\t%.2f\t\t%v\n", department, firstName, lastName, percentage, pass)
}

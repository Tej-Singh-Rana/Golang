package main

import (
	"fmt"
)

func main() {
	// Main code block
	// a slice of type string
  // create a slice from an array, specified size
	var student = [5]string{
		"john",
		"jacob",
		"rahul",
		"raghu",
		"narayan",
	}
	x := student[0:3]
	y := student[0:]
	z := student[2:4]
	final := z[2:3]

	fmt.Println("######### x ########")
	fmt.Println(x)
	fmt.Println("######### y ########")
	fmt.Println(y)
	fmt.Println("######### z ########")
	fmt.Println(z)
	fmt.Println("######### final ########")
	fmt.Println(final)
}


/*
_Output_:-
######### x ########
[john jacob rahul]
######### y ########
[john jacob rahul raghu narayan]
######### z ########
[rahul raghu]
######### final ########
[narayan]

*/

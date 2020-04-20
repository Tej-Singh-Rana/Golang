// specifying entry point of the application

package main

// import block

import "fmt"

func main() {

	// type type-name struct{}
	type user struct {
		ID        int
		firstName string
		lastName  string
	}
	// w/o initializing value it gives null value.
	var student user
	student.ID = 150
	student.firstName = "storm"
	student.lastName = "rider"
	fmt.Println(student)
	fmt.Println(student.lastName)

	trainer := user{ID: 701,
		firstName: "John",
		lastName:  "wick",   // if we will not use comma here then compiler itself mark semicolon(;) here and gives an error. So use comma or }.
	}
	fmt.Println(trainer)

}
